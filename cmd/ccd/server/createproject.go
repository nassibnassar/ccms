package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/log"
)

func createProjectStmt(s *svr, rqid int64, cmd *ast.CreateProjectStmt) *ccms.Result {
	if s.cat.ProjectExists(cmd.ProjectName) {
		return cmderr("project \"" + cmd.ProjectName + "\" already exists")
	}

	if !s.cat.IsValidTargetProject(cmd.ProjectName) {
		return cmderr("invalid project name \"" + cmd.ProjectName + "\"")
	}

	if err := s.cat.CreateProject(cmd.ProjectName); err != nil {
		log.Info("[%d] %v", rqid, err)
		return cmderr("error writing project \"" + cmd.ProjectName + "\"")
	}

	return ccms.NewResult("create project")
}
