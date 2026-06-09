package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/set"
)

func dropSetStmt(s *svr, rqid int64, cmd *ast.DropSetStmt) *ccms.Result {
	set := set.Parse(cmd.Set)

	validTargetSet, err := cat.IsValidTargetSet(s.d, set)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.Set + "\"")
	}

	setExists, err := cat.SetExists(s.d, set)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !setExists {
		return cmderr("set \"" + cmd.Set + "\" does not exist")
	}

	if err := cat.DropSet(s.d, set); err != nil {
		return cmderrint("dropping set", err)
	}

	return ccms.NewResult("drop set")
}
