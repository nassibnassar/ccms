package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func createFilterStmt(s *svr, rqid int64, cmd *ast.CreateFilterStmt) *ccms.Result {
	// TODO validate cmd.Filter

	// TODO error if cmd.Filter is nil

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	_ = sql
	return cmderr("create filter not yet supported")
	// return ccms.NewResult("create filter")
}
