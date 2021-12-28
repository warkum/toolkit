package verihubs

import (
	"time"
)

// Config data struct for verihubs config
type Config struct {
	Host    string        `json:"host" yaml:"host"`
	AppID   string        `json:"app-id" yaml:"app-id"`
	APIKey  string        `json:"api-key" yaml:"api-key"`
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}
