package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
)

func setCounter(ctx context.Context, name string, incr int64, attrs ...attribute.KeyValue) {
	counter, ok := Counters[name]
	if !ok {
		return
	}
	counter.Add(ctx, incr, attrs...)
}

func setUpDownCounter(ctx context.Context, name string, incr int64, attrs ...attribute.KeyValue) {
	upDownCounter, ok := UpDownCounters[name]
	if !ok {
		return
	}
	upDownCounter.Add(ctx, incr, attrs...)
}

func setHistogram(ctx context.Context, name string, incr int64, attrs ...attribute.KeyValue) {
	histogram, ok := Histograms[name]
	if !ok {
		return
	}
	histogram.Record(ctx, incr, attrs...)
}

func SetMetrics(ctx context.Context, name string, metricType Type, incr int64, attrs ...attribute.KeyValue) {
	name = joinMetricName(name, metricType)

	switch metricType {
	case TypeCounter:
		setCounter(ctx, name, incr, attrs...)
	case TypeUpDownCounter:
		setUpDownCounter(ctx, name, incr, attrs...)
	case TypeHistogram:
		setHistogram(ctx, name, incr, attrs...)
	}
}
