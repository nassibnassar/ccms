package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/internal/pgerr"
)

func createFilterStmt(s *svr, rqid int64, cmd *ast.CreateFilterStmt) *ccms.Result {
	// TODO validate cmd.Filter and check if it already exists

	if !cmd.Where.(*ast.WhereClause).Valid {
		return cmderr("required \"where\" clause is missing")
	}

	sql, err := cmd.SQL(s.d)
	if err != nil {
		return cmderr(err.Error())
	}
	q := "insert into ccms.filter (name, sql) values ($1, $2)"
	if _, err := s.d.Q.Exec(s.d.C, q, cmd.Filter, sql); err != nil {
		return cmderr(pgerr.String(err))
	}
	return ccms.NewResult("create filter")
}
