package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func archiveProjectStmt(s *svr, rqid int64, cmd *ast.ArchiveProjectStmt) *ccms.Result {
	if !cat.IsValidTargetProject(cmd.Project) {
		return cmderr("invalid target project \"" + cmd.Project + "\"")
	}

	projectID, err := cat.ProjectID(s.d, cmd.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID == 0 {
		return cmderr("project \"" + cmd.Project + "\" does not exist")
	}
	if projectID == -1 {
		return cmderr("project \"" + cmd.Project + "\" is already archived")
	}

	newProjectName, err := cat.ArchiveProject(s.d, cmd.Project)
	if err != nil {
		return cmderr("archiving project: " + err.Error())
	}

	return ccms.NewResult("archive project " + newProjectName)
}
