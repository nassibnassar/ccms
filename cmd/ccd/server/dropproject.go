package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
)

func dropProjectStmt(s *svr, rqid int64, cmd *ast.DropProjectStmt) *ccms.Result {
	if !cat.IsValidTargetProject(cmd.Project) {
		return cmderr("invalid target project \"" + cmd.Project + "\"")
	}

	projectID, err := cat.ProjectID(s.d, cmd.Project)
	if err != nil {
		return cmderrint("checking if project exists", err)
	}
	if projectID == 0 {
		return cmderr("project \"" + cmd.Project + "\" does not exist")
	}

	if err := cat.DropProject(s.d, cmd.Project); err != nil {
		return cmderrint("dropping project", err)
	}

	return ccms.NewResult("drop project")
}
