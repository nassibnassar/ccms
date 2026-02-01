package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/protocol"
)

func selectStmt(s *svr, rqid int64, cmd *ast.SelectStmt) *protocol.CommandResponse {
	if cmd.Retrieve {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "\"retrieve\" is no longer supported; use \"select\"",
		}
	}

	if !strings.ContainsRune(cmd.From, '.') {
		cmd.From = "ccms." + cmd.From
	}
	if !s.cat.SetExists(cmd.From) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "set \"" + cmd.From + "\" does not exist",
		}
	}

	switch cmd.Select.(type) {
	case *ast.AttrSelectExpr:
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "selecting attributes is not yet supported",
		}
	case *ast.StarSelectExpr:
	}

	if cmd.WhereAttr != "" && !catalog.IsAttribute(cmd.WhereAttr) {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "attribute \"" + cmd.WhereAttr + "\" does not exist",
		}
	}

	switch l := cmd.Limit.(type) {
	case *ast.NoLimitExpr:
		cmd.Limit = &ast.LimitValueExpr{Value: "30"} // temporary maximum
	case *ast.LimitValueExpr:
		lim, _ := strconv.Atoi(l.Value)
		if lim < 0 {
			return &protocol.CommandResponse{
				Status:  "error",
				Message: "limit must not be negative",
			}
		}
		if lim > 30 {
			cmd.Limit = &ast.LimitValueExpr{Value: "30"} // temporary maximum
		}
	}

	log.Info("[%d] %s", rqid, cmd.SQL())
	sql := cmd.SQL()
	rows, err := s.dp.Query(context.TODO(), sql)
	if err != nil {
		fmt.Println(sql)
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
		ids := strconv.FormatInt(id, 10)
		data = append(data, protocol.DataRow{
			Values: []string{ids, author, title, full_vendor_name, availability},
		})
	}
	if err = rows.Err(); err != nil {
		panic(fmt.Sprintf("reading from reserve: %v", err))
	}

	return &protocol.CommandResponse{
		Status: "select",
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
