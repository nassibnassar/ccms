package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/internal/set"
)

func deleteStmt(s *svr, d *dbx.DB, rqid int64, cmd *ast.DeleteStmt) *ccms.Result {
	fromSet := set.Parse(cmd.From)

	validTargetSet, err := cat.IsValidTargetSet(d, fromSet)
	if err != nil {
		return cmderr("checking if target set valid: " + err.Error())
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.From + "\"")
	}

	projectID, err := cat.ProjectID(d, fromSet.Project)
	if err != nil {
		return cmderr("checking if project exists: " + err.Error())
	}
	if projectID == 0 {
		return cmderr("project \"" + fromSet.Project + "\" does not exist")
	}
	if projectID == -1 {
		return cmderr("project \"" + fromSet.Project + "\" is archived")
	}

	setExists, err := cat.SetExists(d, fromSet)
	if err != nil {
		return cmderr("checking if set exists: " + err.Error())
	}
	if !setExists {
		return cmderr("set \"" + cmd.From + "\" does not exist")
	}

	sql, err := cmd.SQL(d)
	if err != nil {
		return cmderr(err.Error())
	}
	if _, err := d.Q.Exec(d.C, sql); err != nil {
		return cmderr("deleting: " + pgerr.String(err))
	}

	return ccms.NewResult("delete")
}
