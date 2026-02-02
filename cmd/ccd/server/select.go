package server

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/indexdata/ccms/cmd/ccd/log"
	"github.com/indexdata/ccms/internal/protocol"
)

func selectStmt(s *svr, rqid int64, cmd *ast.SelectStmt) *protocol.CommandResponse {
	switch cmd.Select.(type) {
	case *ast.AttrSelectExpr:
		return &protocol.CommandResponse{
			Status:  "error",
			Message: "selecting attributes is not yet supported",
		}
	case *ast.StarSelectExpr:
	}

	if err := processQuery(s, rqid, cmd.Query); err != nil {
		return &protocol.CommandResponse{
			Status:  "error",
			Message: err.Error(),
		}
	}

	switch l := cmd.Query.Limit.(type) {
	case *ast.NoLimitExpr:
		cmd.Query.Limit = &ast.LimitValueExpr{Value: "30"} // temporary maximum
	case *ast.LimitValueExpr:
		lim, _ := strconv.Atoi(l.Value)
		if lim < 0 {
			return &protocol.CommandResponse{
				Status:  "error",
				Message: "limit must not be negative",
			}
		}
		if lim > 30 {
			cmd.Query.Limit = &ast.LimitValueExpr{Value: "30"} // temporary maximum
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

func processQuery(s *svr, rqid int64, query *ast.QueryExpr) error {
	if !strings.ContainsRune(query.From, '.') {
		query.From = "ccms." + query.From
	}
	if !s.cat.SetExists(query.From) {
		return errors.New("set \"" + query.From + "\" does not exist")
	}

	if query.WhereAttr != "" && !catalog.IsAttribute(query.WhereAttr) {
		return errors.New("attribute \"" + query.WhereAttr + "\" does not exist")
	}
	return nil
}
