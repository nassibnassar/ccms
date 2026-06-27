package dbx

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

var _ Queryable = (*pgxpool.Pool)(nil)
var _ Queryable = (*pgx.Conn)(nil)
var _ Queryable = (pgx.Tx)(nil)
