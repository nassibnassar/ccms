package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func createProjectStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.CreateProjectStmt) *ccms.Result {
	projectID, err := cat.ProjectID(d, cmd.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID != 0 {
		return cmderr("project \"" + cmd.Project + "\" already exists")
	}

	if !cat.IsValidTargetProject(cmd.Project) {
		return cmderr("invalid project name \"" + cmd.Project + "\"")
	}

	if err := cat.CreateProject(d, cmd.Project); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("create project")
}
