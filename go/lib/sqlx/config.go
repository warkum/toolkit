package sqlx

import "time"

const defaultMaxRetry = 1

// Config database connection configuration
type Config struct {
	Driver      string `json:"driver" yaml:"driver"`
	MasterDSN   string `json:"master" yaml:"master"`
	FollowerDSN string `json:"follower" yaml:"follower"`

	// number of retry during Connect
	// won't be used if `NoPingOnOpen`=true
	Retry int `json:"retry" yaml:"retry"`

	// use config that specific for each master/slave
	MaxOpenConnections    int           `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConnections    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnectionMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	// no Ping when opening DB connection, useful if we don't care whether the server is up or not
	NoPingOnOpen bool `json:"no_ping_on_open" yaml:"no_ping_on_open"`
}
