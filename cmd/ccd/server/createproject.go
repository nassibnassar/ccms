package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func createProjectStmt(s *svr, rqid int64, cmd *ast.CreateProjectStmt) *ccms.Result {
	projectID, err := cat.ProjectID(s.d, cmd.Project)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if projectID != 0 {
		return cmderr("project \"" + cmd.Project + "\" already exists")
	}

	if !cat.IsValidTargetProject(cmd.Project) {
		return cmderr("invalid project name \"" + cmd.Project + "\"")
	}

	if err := cat.CreateProject(s.d, cmd.Project); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("create project")
}
