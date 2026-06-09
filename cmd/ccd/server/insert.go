package server

import (
	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/set"
)

func insertStmt(s *svr, rqid int64, cmd *ast.InsertStmt) *ccms.Result {
	o := cmd.Query.(*ast.QueryClause).Order.(*ast.OrderClause)
	if o.Valid {
		return cmderr("\"order by\" is not supported with insert")
	}
	f := cmd.Query.(*ast.QueryClause).Offset.(*ast.OffsetClause)
	if f.Valid {
		return cmderr("\"offset\" is not supported with insert")
	}

	intoSet := set.Parse(cmd.Into)
	validTargetSet, err := cat.IsValidTargetSet(s.d, intoSet)
	if err != nil {
		return cmderrint("checking if target set valid", err)
	}
	if !validTargetSet {
		return cmderr("invalid target set \"" + cmd.Into + "\"")
	}

	intoSetExists, err := cat.SetExists(s.d, intoSet)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !intoSetExists {
		return cmderr("set \"" + cmd.Into + "\" does not exist")
	}

	from := cmd.Query.(*ast.QueryClause).From
	if from == "reserve" { // TODO remove this "reserve" check after some time
		return cmderr("set \"reserve\" is no longer supported; use \"<project>.object\"")
	}
	fromSet := set.Parse(from)
	fromSetExists, err := cat.SetExists(s.d, fromSet)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !fromSetExists {
		return cmderr("set \"" + from + "\" does not exist")
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	if _, err := s.d.Q.Exec(s.d.C, sql); err != nil {
		return cmderrint("inserting data into \""+cmd.Into+"\"", err)
	}

	return ccms.NewResult("insert")
}
