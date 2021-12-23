package sqlx

import (
	"github.com/jmoiron/sqlx"
)

// Package dependencies for sqlx package
type Package struct {
	master   *sqlx.DB
	follower *sqlx.DB
}

// New create sqlx connection using config
func New(cfg Config) (*Package, error) {
	masterOpen, err := openConnection(cfg.MasterDSN, cfg)
	if err != nil {
		return nil, err
	}

	followerOpen, err := openConnection(cfg.FollowerDSN, cfg)
	if err != nil {
		return nil, err
	}

	pkg := &Package{
		master:   masterOpen,
		follower: followerOpen,
	}

	return pkg, nil
}

// openConnection opening DB connection by using config
// will be returning sqlx DB
func openConnection(dsn string, cfg Config) (*sqlx.DB, error) {
	if cfg.Retry < 1 || cfg.NoPingOnOpen {
		cfg.Retry = defaultMaxRetry
	}

	var (
		db  *sqlx.DB
		err error
	)

	for i := 1; i <= cfg.Retry; i++ {
		db, err = sqlx.Open(cfg.Driver, dsn)
		if err != nil {
			continue
		}

		break
	}

	if err != nil {
		return nil, err
	}

	if !cfg.NoPingOnOpen {
		err = db.Ping()
		if err != nil {
			_ = db.Close()
			return nil, err
		}
	}

	if cfg.ConnectionMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)
	}

	if cfg.MaxOpenConnections > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConnections)
	}

	if cfg.MaxIdleConnections > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
	}

	return db, nil
}
