package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/internal/protocol"
)

func selectStmt(s *svr, cmd *ast.SelectStmt) *protocol.CommandResponse {
	if cmd.Retrieve {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "\"retrieve all\" is no longer supported; use \"select *\"",
		}
	}

	a, ok := cmd.Select.(*ast.AttrSelectExpr)
	if ok {
		if a.Attribute != "all" {
			return &protocol.CommandResponse{
				Status:  "error",
				Message: "selecting attributes is not yet supported",
			}
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
	q := "select id, coalesce(author, ''), coalesce(title, ''), coalesce(full_vendor_name, ''), coalesce(availability, '') from ccms.attr where author is not null limit " + strconv.Itoa(lim)
	rows, err := s.dp.Query(context.TODO(), q)
	if err != nil {
		panic(fmt.Sprintf("selecting from reserve: %v", err))
	}
	defer rows.Close()
	data := make([]protocol.DataRow, 0)
	for rows.Next() {
		var id int64
		var author, title, full_vendor_name, availability string
		err = rows.Scan(&id, &author, &title, &full_vendor_name, &availability)
		if err != nil {
			panic(fmt.Sprintf("reading from reserve: %v", err))
		}
		data = append(data, protocol.DataRow{
			Values: []string{strconv.FormatInt(id, 10), author, title, full_vendor_name, availability},
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
			{
				Name: "full_vendor_name",
			},
			{
				Name: "availability",
			},
		},
		Data: data,
	}
}
