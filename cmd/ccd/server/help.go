package server

import (
	"fmt"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/doc"
	"github.com/indexdata/ccms/internal/protocol"
)

func helpStmt(s *svr, cmd *ast.HelpStmt) *protocol.CommandResponse {
	if cmd.Topic == "" {
		return &protocol.CommandResponse{
			Status: "help",
			Fields: []protocol.FieldDescription{
				{
					Name: "help",
				},
			},
			Data: []protocol.DataRow{
				{
					Values: []string{`select        retrieve objects from a set
show filters  list existing filters
show sets     list existing sets`},
				},
			},
		}
	}

	var docstr string
	switch cmd.Topic {
	case "select":
		docstr = doc.Select()
	case "show filters":
		docstr = doc.ShowFilters()
	case "show sets":
		docstr = doc.ShowSets()
	default:
		docstr = fmt.Sprintf("unknown command %q", cmd.Topic)
	}

	return &protocol.CommandResponse{
		Status: "help",
		Fields: []protocol.FieldDescription{
			{
				Name: "help",
			},
		},
		Data: []protocol.DataRow{
			{
				Values: []string{docstr},
			},
		},
	}

	//if cmd.Retrieve {
	//        return &protocol.CommandResponse{
	//                Status:  "error",
	//                Message: "\"retrieve all\" is no longer supported; use \"select *\"",
	//        }
	//}

}
