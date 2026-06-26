package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func dropProjectStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.DropProjectStmt) *ccms.Result {
	if !cat.IsValidTargetProject(cmd.Project) {
		return cmderr("invalid target project \"" + cmd.Project + "\"")
	}

	projectID, err := cat.ProjectID(d, cmd.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID == 0 {
		return cmderr("project \"" + cmd.Project + "\" does not exist")
	}
	if projectID != -1 {
		return cmderr("project \"" + cmd.Project + "\" is not archived")
	}

	if err := cat.DropProject(d, cmd.Project); err != nil {
		return cmderr("dropping project: " + err.Error())
	}

	return ccms.NewResult("drop project")
}
