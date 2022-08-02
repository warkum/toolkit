package metrics

import (
	"go.opentelemetry.io/otel/metric/instrument"
)

func registerCounter(name string, opts ...instrument.Option) error {
	_, ok := Counters[name]
	if ok {
		return nil
	}

	counterNew, err := meter.SyncInt64().Counter(name, opts...)
	if err != nil {
		return err
	}

	Counters[name] = counterNew

	return nil
}

func registerUpDownCounter(name string, opts ...instrument.Option) error {
	_, ok := UpDownCounters[name]
	if ok {
		return nil
	}

	upDownCounterNew, err := meter.SyncInt64().UpDownCounter(name, opts...)
	if err != nil {
		return err
	}

	UpDownCounters[name] = upDownCounterNew

	return nil
}

func registerHistogram(name string, opts ...instrument.Option) error {
	_, ok := Histograms[name]
	if ok {
		return nil
	}

	histogramNew, err := meter.SyncInt64().Histogram(name, opts...)
	if err != nil {
		return err
	}

	Histograms[name] = histogramNew

	return nil
}

func RegisterMetrics(name string, metricType Type, opts ...instrument.Option) error {
	if meter == nil {
		return ErrMetersNotInitialized
	}

	name = joinMetricName(name, metricType)

	switch metricType {
	case TypeCounter:
		return registerCounter(name, opts...)
	case TypeUpDownCounter:
		return registerUpDownCounter(name, opts...)
	case TypeHistogram:
		return registerHistogram(name, opts...)
	}

	return ErrInvalidType
}

func RegisterUsingConfigs(configs []Config) error {
	if meter == nil {
		return ErrMetersNotInitialized
	}

	for _, config := range configs {
		opts := make([]instrument.Option, 0, 2)

		if config.Description != "" {
			opts = append(opts, instrument.WithDescription(config.Description))
		}

		if config.Unit != "" {
			opts = append(opts, instrument.WithUnit(config.Unit))
		}

		err := RegisterMetrics(config.Name, config.Type, opts...)
		if err != nil {
			return err
		}
	}

	return nil
}
