package server

import (
	"fmt"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/doc"
	"github.com/indexdata/ccms/internal/protocol"
)

func infoStmt(s *svr, cmd *ast.InfoStmt) *protocol.CommandResponse {
	if cmd.Topic == "" {
		return &protocol.CommandResponse{
			Status: "info",
			Fields: []protocol.FieldDescription{
				{
					Name: "info",
				},
			},
			Data: []protocol.DataRow{
				{
					Values: []string{"" +
						"SQL commands:\n" +
						"        create set  define a new set\n" +
						//"        info    show supported commands\n" +
						"        insert      insert objects into a set\n" +
						"        select      retrieve objects from a set\n" +
						"        show        list existing filters or sets\n"},
				},
			},
		}
	}

	var docstr string
	switch cmd.Topic {
	case "create set":
		docstr = doc.CreateSet()
	//case "info":
	//        docstr = doc.Info()
	case "insert":
		docstr = doc.Insert()
	case "select":
		docstr = doc.Select()
	case "show":
		docstr = doc.Show()
	default:
		docstr = fmt.Sprintf("unknown command %q", cmd.Topic)
	}

	return &protocol.CommandResponse{
		Status: "info",
		Fields: []protocol.FieldDescription{
			{
				Name: "info",
			},
		},
		Data: []protocol.DataRow{
			{
				Values: []string{docstr},
			},
		},
	}
}
