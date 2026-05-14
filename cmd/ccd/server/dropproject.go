package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
)

func dropProjectStmt(s *svr, rqid int64, cmd *ast.DropProjectStmt) *ccms.Result {
	if !s.cat.IsValidTargetProject(cmd.ProjectName) {
		return cmderr("invalid target project \"" + cmd.ProjectName + "\"")
	}

	if !s.cat.ProjectExists(cmd.ProjectName) {
		return cmderr("project \"" + cmd.ProjectName + "\" does not exist")
	}

	if err := s.cat.DropProject(cmd.ProjectName); err != nil {
		return cmderr(err.Error())
	}

	return ccms.NewResult("drop project")
}
