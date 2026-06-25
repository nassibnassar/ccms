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
	e, ok := err.(*pgconn.PgError)
	if !ok {
		// log.Info("internal server error: pgerr: error is type %T", err)
		return err.Error()
	}
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
