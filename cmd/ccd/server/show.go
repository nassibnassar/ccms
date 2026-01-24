package server

import (
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/internal/protocol"
)

func showStmt(s *svr, cmd *ast.ShowStmt) *protocol.CommandResponse {
	switch cmd.Name {
	case "filters":
		return &protocol.CommandResponse{
			Status: "show",
			Fields: []protocol.FieldDescription{
				{
					Name: "filter",
				},
			},
			Data: []protocol.DataRow{},
		}
	case "sets":
		return &protocol.CommandResponse{
			Status: "show sets",
			Fields: []protocol.FieldDescription{
				{
					Name: "set",
				},
			},
			Data: []protocol.DataRow{
				{
					Values: []string{"reserve"},
				},
			},
		}
	default:
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "unknown variable \"" + cmd.Name + "\"",
		}
	}
}
