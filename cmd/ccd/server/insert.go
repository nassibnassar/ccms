package server

import (
	"context"
	"fmt"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/protocol"
)

func insertStmt(s *svr, rqid int64, cmd *ast.InsertStmt) *protocol.CommandResponse {
	if err := processQuery(s, rqid, cmd.Query); err != nil {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: err.Error(),
		}
	}

	log.Info("[%d] %s", rqid, cmd.SQL())
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
