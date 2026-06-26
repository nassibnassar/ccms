package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/cmd/ccd/dbx"
)

func alterProjectStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.AlterProjectStmt) *ccms.Result {
	projectID, err := cat.ProjectID(d, cmd.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID == 0 {
		return cmderr("project \"" + cmd.Project + "\" does not exist")
	}
	if projectID == -1 {
		return cmderr("project \"" + cmd.Project + "\" is archived")
	}

	switch cmd.Action {
	case ast.Set:
		if err := cat.AlterProjectSetProperty(d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Add:
		if err := cat.AlterProjectAddToProperty(d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Drop:
		if err := cat.AlterProjectDropFromProperty(d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	default:
		return cmderr(internalError + "unknown action in alter project")
	}

	return ccms.NewResult("alter project")
}
