package metrics

import "go.opentelemetry.io/otel/metric/unit"

// Config detail struct for metric config
type Config struct {
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	Unit        unit.Unit `json:"unit" yaml:"unit"` //1, By or ms
	Type        Type      `json:"type" yaml:"type"`
}

// InitConfig detail struct for init config
type InitConfig struct {
	AppName string   `json:"app_name" yaml:"app_name"`
	Configs []Config `json:"configs" yaml:"configs"`
}
