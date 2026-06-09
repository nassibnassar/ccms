package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func alterProjectStmt(s *svr, rqid int64, cmd *ast.AlterProjectStmt) *ccms.Result {
	projectID, err := cat.ProjectID(s.d, cmd.Project)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if projectID == 0 {
		return cmderr("project \"" + cmd.Project + "\" does not exist")
	}

	switch cmd.Action {
	case ast.Set:
		if err := cat.AlterProjectSetProperty(s.d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Add:
		if err := cat.AlterProjectAddToProperty(s.d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Drop:
		if err := cat.AlterProjectDropFromProperty(s.d, cmd.Project, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	default:
		return cmderr(internalError + "unknown action in alter project")
	}

	return ccms.NewResult("alter project")
}
