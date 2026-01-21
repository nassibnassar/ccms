package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/internal/protocol"
)

//	type RetrieveStmt struct {
//	       Attribute string
//	       Set       string
//	       Limit     int
//	}
func retrieve(s *svr, cmd *ast.RetrieveStmt) *protocol.CommandResponse {
	if cmd.Attribute != "all" {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "selecting attributes not yet supported; use \"retrieve all\"",
		}
	}
	if cmd.Set != "reserve" {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.Set + "\" does not exist",
		}
	}
	lim, _ := strconv.Atoi(cmd.Limit)
	if lim > 20 {
		lim = 20
	}
	q := "select id, coalesce(author, ''), coalesce(title, '') from ccms.attr where author is not null limit " + strconv.Itoa(lim)
	rows, err := s.dp.Query(context.TODO(), q)
	if err != nil {
		panic(fmt.Sprintf("selecting from reserve: %v", err))
	}
	defer rows.Close()
	data := make([]protocol.DataRow, 0)
	for rows.Next() {
		var id int64
		var author, title string
		err = rows.Scan(&id, &author, &title)
		if err != nil {
			panic(fmt.Sprintf("reading from reserve: %v", err))
		}
		data = append(data, protocol.DataRow{
			Values: []string{strconv.FormatInt(id, 10), author, title},
		})
	}
	if err = rows.Err(); err != nil {
		panic(fmt.Sprintf("reading from reserve: %v", err))
	}

	return &protocol.CommandResponse{
		Status: "retrieve",
		Fields: []protocol.FieldDescription{
			{
				Name: "id",
			},
			{
				Name: "author",
			},
			{
				Name: "title",
			},
		},
		Data: data,
	}
}
