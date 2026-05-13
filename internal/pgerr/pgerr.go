package pgerr

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

func Error(err error) error {
	return errors.New(String(err))
}

func String(err error) string {
	e := err.(*pgconn.PgError)
	var b strings.Builder
	b.WriteString(e.Message)
	if e.Detail != "" {
		b.WriteString(": ")
		b.WriteString(e.Detail)
	}
	if e.Hint != "" {
		b.WriteString(" (")
		b.WriteString(e.Hint)
		b.WriteRune(')')
	}
	return b.String()
}
