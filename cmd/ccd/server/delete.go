package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/pgerr"
)

func deleteStmt(s *svr, rqid int64, cmd *ast.DeleteStmt) *ccms.Result {
	validTargetSet, err := cat.IsValidTargetSet(s.d, cmd.From)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.From + "\"")
	}

	setExists, err := cat.SetExists(s.d, cmd.From)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !setExists {
		return cmderr("set \"" + cmd.From + "\" does not exist")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	if _, err := s.d.Q.Exec(s.d.C, sql); err != nil {
		return cmderrint("deleting", pgerr.Error(err))
	}

	return ccms.NewResult("delete")
}
