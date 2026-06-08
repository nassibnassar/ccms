package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func alterProjectStmt(s *svr, rqid int64, cmd *ast.AlterProjectStmt) *ccms.Result {
	projectExists, err := cat.ProjectExists(s.d, cmd.ProjectName)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if !projectExists {
		return cmderr("project \"" + cmd.ProjectName + "\" does not exist")
	}

	switch cmd.Action {
	case ast.Set:
		if err := cat.AlterProjectSetProperty(s.d, cmd.ProjectName, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Add:
		if err := cat.AlterProjectAddToProperty(s.d, cmd.ProjectName, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	case ast.Drop:
		if err := cat.AlterProjectDropFromProperty(s.d, cmd.ProjectName, cmd.Property, cmd.Value, cmd.StringLiteral); err != nil {
			return cmderr(err.Error())
		}
	default:
		return cmderr(internalError + "unknown action in alter project")
	}

	return ccms.NewResult("alter project")
}
