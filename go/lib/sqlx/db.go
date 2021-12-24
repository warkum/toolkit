package sqlx

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// GetMasterDB get master DB connection
func (p *Package) GetMasterDB() *sqlx.DB {
	return p.master
}

// GetFollowerDB get follower DB connection
func (p *Package) GetFollowerDB() *sqlx.DB {
	return p.follower
}

// TxMaster begin transaction using master DB connection
func (p *Package) TxMaster() (*sqlx.Tx, error) {
	return p.master.Beginx()
}

// TxFollower begin transaction using follower DB connection
func (p *Package) TxFollower() (*sqlx.Tx, error) {
	return p.follower.Beginx()
}

// Rollbackx rollback transaction without returning error
func (p *Package) Rollbackx(tx *sqlx.Tx) {
	if tx == nil {
		return
	}

	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		logrus.Warn(err)
	}
}
