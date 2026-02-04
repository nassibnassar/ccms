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
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "invalid target set \"" + cmd.Into + "\"",
		}
	}

	if !s.cat.SetExists(cmd.Into) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.Into + "\" does not exist",
		}
	}

	if err := processQuery(s, rqid, cmd.Query); err != nil {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: err.Error(),
		}
	}

	//log.Info("[%d] %s", rqid, cmd.SQL())
	sql := cmd.SQL()
	if _, err := s.dp.Exec(context.TODO(), sql); err != nil {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: fmt.Sprintf("inserting data into %q: %v", cmd.Into, err),
		}
	}

	return &protocol.CommandResponse{
		Status: "insert",
	}
}
