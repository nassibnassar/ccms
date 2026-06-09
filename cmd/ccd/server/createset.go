package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/set"
)

func createSetStmt(s *svr, rqid int64, cmd *ast.CreateSetStmt) *ccms.Result {
	set := set.Parse(cmd.Set)

	setExists, err := cat.SetExists(s.d, set)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if setExists {
		return cmderr("set \"" + cmd.Set + "\" already exists")
	}

	validTargetSet, err := cat.IsValidTargetSet(s.d, set)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid set name \"" + cmd.Set + "\"")
	}

	projectExists, err := cat.ProjectExists(s.d, set.Project)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if !projectExists {
		return cmderr("invalid project in  \"" + cmd.Set + "\"")
	}

	if err := cat.CreateSet(s.d, set); err != nil {
		return cmderrint("writing set", err)
	}

	return ccms.NewResult("create set")
}
