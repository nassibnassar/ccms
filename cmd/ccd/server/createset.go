package server

import (
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/protocol"
)

func createSetStmt(s *svr, rqid int64, cmd *ast.CreateSetStmt) *protocol.CommandResponse {
	// TODO check if project exists

	if cmd.SetName == "reserve" || s.cat.SetExists(cmd.SetName) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.SetName + "\" already exists",
		}
	}

	if !strings.ContainsRune(cmd.SetName, '.') {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.SetName + "\" does not have a project name",
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
