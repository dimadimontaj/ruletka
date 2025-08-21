package cmd

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSqlConn(configuration *postgresConfiguration) (*sql.DB, error) {
	db, err := sql.Open("postgres", configuration.GetConnectionString())
	if err != nil {
		return nil, errors.Wrap(err, "cant connect to db")
	}

	db.SetMaxIdleConns(configuration.GetMaxIdleConns())
	db.SetMaxOpenConns(configuration.GetMaxOpenConns())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "cant ping db")
	}

	return db, nil
}