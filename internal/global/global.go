package global

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

// Version is defined at build time via -ldflags.
var Version = ""

const DefaultPort = "8504"

const ServerProgram = "ccd"

const ClientProgram = "ccc"

func ServerConfigFileName(datadir string) string {
	return filepath.Join(datadir, "ccd.conf")
}

func PGErr(err error) error {
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
	return errors.New(b.String())
}
