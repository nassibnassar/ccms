package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/internal/set"
)

func deleteStmt(s *svr, rqid int64, cmd *ast.DeleteStmt) *ccms.Result {
	fromSet := set.Parse(cmd.From)

	validTargetSet, err := cat.IsValidTargetSet(s.d, fromSet)
	if err != nil {
		return cmderr("checking if target set valid: " + err.Error())
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.From + "\"")
	}

	projectID, err := cat.ProjectID(s.d, fromSet.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID == 0 {
		return cmderr("project \"" + fromSet.Project + "\" does not exist")
	}
	if projectID == -1 {
		return cmderr("project \"" + fromSet.Project + "\" is archived")
	}

	setExists, err := cat.SetExists(s.d, fromSet)
	if err != nil {
		return cmderr("checking if set exists: " + err.Error())
	}
	if !setExists {
		return cmderr("set \"" + cmd.From + "\" does not exist")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	if _, err := s.d.Q.Exec(s.d.C, sql); err != nil {
		return cmderr("deleting: " + pgerr.Error(err).Error())
	}

	return ccms.NewResult("delete")
}
