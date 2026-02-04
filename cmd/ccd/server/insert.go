package server

import (
	"context"
	"fmt"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/internal/protocol"
)

func insertStmt(s *svr, rqid int64, cmd *ast.InsertStmt) *protocol.CommandResponse {
	if !catalog.IsValidTargetSet(cmd.Into) {
		return cmderr("invalid target set \"" + cmd.Into + "\"")
	}

	if !s.cat.SetExists(cmd.Into) {
		return cmderr("set \"" + cmd.Into + "\" does not exist")
	}

	if err := processQuery(s, rqid, cmd.Query); err != nil {
		return cmderr(err.Error())
	}

	switch cmd.Query.Order.(type) {
	case *ast.OrderValueExpr:
		return cmderr("\"order by\" is not supported with insert")
	}

	//log.Info("[%d] %s", rqid, cmd.SQL())
	sql := cmd.SQL()
	if _, err := s.dp.Exec(context.TODO(), sql); err != nil {
		return cmderr(fmt.Sprintf("inserting data into %q: %v", cmd.Into, err))
	}

	return &protocol.CommandResponse{
		Status: "insert",
	}
}
