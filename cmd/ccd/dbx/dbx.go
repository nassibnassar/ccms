package dbx

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func Connect(ctx context.Context, connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// func newDBConn(ctx context.Context, connString string) (*pgx.Conn, error) {
// 	config, err := pgx.ParseConfig(connString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// config.AfterConnect = setDatabaseParameters
// 	// config.MaxConns = 64
// 	dc, err := pgx.ConnectConfig(ctx, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return dc, nil
// }

// func newPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
// 	config, err := pgxpool.ParseConfig(connString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	config.AfterConnect = setDatabaseParameters
// 	config.MaxConns = 64
// 	dp, err := pgxpool.NewWithConfig(ctx, config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return dp, nil
// }

// func setDatabaseParameters(ctx context.Context, conn *pgx.Conn) error {
// 	q := "set search_path = 'public'"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	q = "set idle_in_transaction_session_timeout=0"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	q = "set idle_session_timeout=0"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	q = "set statement_timeout=0"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	q = "set timezone='UTC'"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	q = "set default_transaction_isolation=serializable"
// 	if _, err := conn.Exec(ctx, q); err != nil {
// 		return dberr.Error(err)
// 	}
// 	return nil
// }

type Table struct {
	Schema string
	Table  string
}

func ParseTable(table string) Table {
	s := strings.Split(table, ".")
	if len(s) == 2 {
		return Table{Schema: s[0], Table: s[1]}
	} else {
		return Table{}
	}
}

func (t Table) String() string {
	return t.Schema + "." + t.Table
}

type DB struct {
	Ctx context.Context
	Queryable
}

type Queryable interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)
}

// var _ Queryable = (*pgxpool.Pool)(nil)
var _ Queryable = (*pgx.Conn)(nil)
var _ Queryable = (pgx.Tx)(nil)
