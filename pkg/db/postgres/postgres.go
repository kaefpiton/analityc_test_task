package postgres

import (
	"analityc_test_task/cmd/config"
	"analityc_test_task/pkg/logger"
	"database/sql"
)
import _ "github.com/lib/pq"

type DB struct {
	*sql.DB
	log logger.Logger
}

func NewDBConnection(cnf *config.Config, log logger.Logger) (*DB, error) {
	db, err := sql.Open("postgres", cnf.GetPgDsn())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cnf.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cnf.Postgres.MaxIdleConns)

	return &DB{
		db,
		log,
	}, nil
}
