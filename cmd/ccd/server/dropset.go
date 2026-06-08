package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func dropSetStmt(s *svr, rqid int64, cmd *ast.DropSetStmt) *ccms.Result {
	validTargetSet, err := cat.IsValidTargetSet(s.d, cmd.SetName)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.SetName + "\"")
	}

	setExists, err := cat.SetExists(s.d, cmd.SetName)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !setExists {
		return cmderr("set \"" + cmd.SetName + "\" does not exist")
	}

	if err := cat.DropSet(s.d, cmd.SetName); err != nil {
		return cmderrint("dropping set", err)
	}

	return ccms.NewResult("drop set")
}
