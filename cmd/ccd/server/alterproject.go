package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func alterProjectStmt(s *svr, rqid int64, cmd *ast.AlterProjectStmt) *ccms.Result {
	if !s.cat.ProjectExists(cmd.ProjectName) {
		return cmderr("project \"" + cmd.ProjectName + "\" does not exist")
	}

	switch cmd.Action {
	case ast.Set:
		if err := s.cat.AlterProjectSetProperty(cmd.ProjectName, cmd.Property, cmd.Value); err != nil {
			return cmderr(err.Error())
		}
	case ast.Add:
		if err := s.cat.AlterProjectAddToProperty(cmd.ProjectName, cmd.Property, cmd.Value); err != nil {
			return cmderr(err.Error())
		}
	case ast.Drop:
		if err := s.cat.AlterProjectDropFromProperty(cmd.ProjectName, cmd.Property, cmd.Value); err != nil {
			return cmderr(err.Error())
		}
	default:
		return cmderr("internal error: unknown action in alter project")
	}

	return ccms.NewResult("alter project")
}
