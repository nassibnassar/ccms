package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createSetStmt(s *svr, rqid int64, cmd *ast.CreateSetStmt) *ccms.Result {
	setExists, err := cat.SetExists(s.d, cmd.SetName)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if setExists {
		return cmderr("set \"" + cmd.SetName + "\" already exists")
	}

	validTargetSet, err := cat.IsValidTargetSet(s.d, cmd.SetName)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid set name \"" + cmd.SetName + "\"")
	}

	projectExists, err := cat.ProjectExists(s.d, cmd.SetName)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if !projectExists {
		return cmderr("invalid project in  \"" + cmd.SetName + "\"")
	}

	if err := cat.CreateSet(s.d, cmd.SetName); err != nil {
		return cmderrint("writing set", err)
	}

	return ccms.NewResult("create set")
}
