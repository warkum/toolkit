package metrics

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var (
	meter metric.Meter

	Counters       = make(map[string]syncint64.Counter)
	UpDownCounters = make(map[string]syncint64.UpDownCounter)
	Histograms     = make(map[string]syncint64.Histogram)
)

func Init(appName string, cfg prometheus.Config) (*prometheus.Exporter, error) {
	res, err := resource.New(context.TODO(),
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(appName),
		))
	if err != nil {
		return nil, err
	}

	ctrl := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(
				histogram.WithExplicitBoundaries(cfg.DefaultHistogramBoundaries),
			),
			aggregation.CumulativeTemporalitySelector(),
			processor.WithMemory(true),
		),
		controller.WithResource(res),
	)

	exporter, err := prometheus.New(cfg, ctrl)
	if err != nil {
		return nil, err
	}

	global.SetMeterProvider(exporter.MeterProvider())

	err = runtime.Start()
	if err != nil {
		return nil, err
	}

	meter = global.MeterProvider().Meter(appName)
	return exporter, nil
}

func InitUsingConfig(cfg InitConfig) (*prometheus.Exporter, error) {
	exporter, err := Init(cfg.AppName, prometheus.Config{})
	if err != nil {
		return nil, err
	}

	err = RegisterUsingConfigs(cfg.Configs)
	return exporter, err
}

func joinMetricName(name string, metricType Type) string {
	return name + NameSeparator + string(metricType)
}
