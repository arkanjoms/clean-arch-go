package database

import (
	"context"
	"database/sql"
)

var ErrNoRows = sql.ErrNoRows

type Database interface {
	Many(ctx context.Context, query string, params ...interface{}) (*sql.Rows, error)
	One(ctx context.Context, query string, params ...interface{}) *sql.Row
	Exec(ctx context.Context, query string, params ...interface{}) (sql.Result, error)
}
