package server

import (
	"strings"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/pgerr"
)

func createFilterStmt(s *svr, rqid int64, cmd *ast.CreateFilterStmt) *ccms.Result {
	filterExists, err := cat.FilterExists(s.d, cmd.Filter)
	if err != nil {
		return cmderr(err.Error())
	}
	if filterExists {
		return cmderr("filter \"" + cmd.Filter + "\" already exists")
	}

	if !cat.IsValidFilterName(cmd.Filter) {
		return cmderr("invalid filter name \"" + cmd.Filter + "\"")
	}

	if !cmd.Where.(*ast.WhereClause).Valid {
		return cmderr("required \"where\" clause is missing")
	}

	var a strings.Builder
	a.WriteString("create filter ")
	a.WriteString(cmd.Filter)
	a.WriteString(" where ")

	sql, err := cmd.SQL(s.d, &a)
	if err != nil {
		return cmderr(err.Error())
	}
	q := "insert into ccms.filter (name, command, sql) values ($1, $2, $3)"
	if _, err := s.d.Q.Exec(s.d.C, q, cmd.Filter, a.String(), sql); err != nil {
		return cmderr(pgerr.String(err))
	}
	return ccms.NewResult("create filter")
}
