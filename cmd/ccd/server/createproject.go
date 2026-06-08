package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createProjectStmt(s *svr, rqid int64, cmd *ast.CreateProjectStmt) *ccms.Result {
	projectExists, err := cat.ProjectExists(s.d, cmd.ProjectName)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if projectExists {
		return cmderr("project \"" + cmd.ProjectName + "\" already exists")
	}

	if !cat.IsValidTargetProject(cmd.ProjectName) {
		return cmderr("invalid project name \"" + cmd.ProjectName + "\"")
	}

	if err := cat.CreateProject(s.d, cmd.ProjectName); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("create project")
}
