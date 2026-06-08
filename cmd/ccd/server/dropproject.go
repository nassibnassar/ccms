package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func dropProjectStmt(s *svr, rqid int64, cmd *ast.DropProjectStmt) *ccms.Result {
	if !cat.IsValidTargetProject(cmd.ProjectName) {
		return cmderr("invalid target project \"" + cmd.ProjectName + "\"")
	}

	projectExists, err := cat.ProjectExists(s.d, cmd.ProjectName)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if !projectExists {
		return cmderr("invalid project in  \"" + cmd.ProjectName + "\"")
	}

	if !projectExists {
		return cmderr("project \"" + cmd.ProjectName + "\" does not exist")
	}

	if err := cat.DropProject(s.d, cmd.ProjectName); err != nil {
		return cmderrint("dropping project", err)
	}

	return ccms.NewResult("drop project")
}
