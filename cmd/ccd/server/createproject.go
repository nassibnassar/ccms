package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func createProjectStmt(s *svr, rqid int64, cmd *ast.CreateProjectStmt) *ccms.Result {
	if s.cat.ProjectExists(cmd.ProjectName) {
		return cmderr("project \"" + cmd.ProjectName + "\" already exists")
	}

	if !s.cat.IsValidTargetProject(cmd.ProjectName) {
		return cmderr("invalid project name \"" + cmd.ProjectName + "\"")
	}

	if err := s.cat.CreateProject(cmd.ProjectName); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("create project")
}
