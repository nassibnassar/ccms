package server

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/protocol"
)

func createSetStmt(s *svr, rqid int64, cmd *ast.CreateSetStmt) *protocol.CommandResponse {
	if s.cat.SetExists(cmd.SetName) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.SetName + "\" already exists",
		}
	}

	if !catalog.IsValidTargetSet(cmd.SetName) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "invalid set name \"" + cmd.SetName + "\"",
		}
	}

	if !catalog.ProjectExists(cmd.SetName) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "invalid project in  \"" + cmd.SetName + "\"",
		}
	}

	if err := s.cat.CreateSet(cmd.SetName); err != nil {
		log.Info("[%d] %v", rqid, err)
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "error writing set \"" + cmd.SetName + "\"",
		}
	}

	return &protocol.CommandResponse{
		Status: "create set",
	}
}
