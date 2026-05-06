package server

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/catalog"
	"github.com/jackc/pgx/v5"
)

func selectStmt(s *svr, rqid int64, cmd *ast.SelectStmt) *ccms.Result {

	f := cmd.Query.(*ast.QueryClause).Offset.(*ast.OffsetClause)
	if f.Valid {
		o := cmd.Query.(*ast.QueryClause).Order.(*ast.OrderClause)
		if !o.Valid {
			return cmderr("\"order by\" is required when \"offset\" is used")
		}
	}

	a := cmd.AttrList.(*ast.SelectAttrList)
	if a.Attr != "*" && a.Attr != "count(*)" {
		return cmderr("selecting attributes is not yet supported")
	}

	if err := processQuery(s, rqid, cmd.Query.(*ast.QueryClause)); err != nil {
		return cmderr(err.Error())
	}

	o := cmd.Query.(*ast.QueryClause).Order.(*ast.OrderClause)
	if o.Valid {
		if !catalog.IsAttr(o.Attr) {
			return cmderr("attribute \"" + o.Attr + "\" does not exist")
		}
	}

	q := cmd.Query.(*ast.QueryClause)
	l := q.Limit.(*ast.LimitClause)
	if l.Valid {
		lim, _ := strconv.Atoi(l.Count)
		if lim < 0 {
			return cmderr("limit must not be negative")
		}
	}

	sql, err := cmd.SQL()
	if err != nil {
		return cmderr(err.Error())
	}
	//log.Info("[%d] %s", rqid, sql)

	switch a.Attr {
	case "*":
		result, err := runQuery(s, sql)
		if err != nil {
			return cmderr(err.Error())
		}
		return result
	case "count(*)":
		result, err := runQueryCount(s, sql)
		if err != nil {
			return cmderr(err.Error())
		}
		return result
	default:
		return cmderr("internal error: invalid projection in select")
	}
}

func runQueryCount(s *svr, sql string) (*ccms.Result, error) {
	var count int64
	err := s.dp.QueryRow(context.TODO(), sql).Scan(&count)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, errors.New("internal error: no rows in select count(*)")
	case err != nil:
		return nil, err
	default:
	}
	result := ccms.NewResult("select")
	result.AddField("count", "bigint")
	result.AddData([]any{count})
	return result, nil
}

func runQuery(s *svr, sql string) (*ccms.Result, error) {
	rows, err := s.dp.Query(context.TODO(), sql)
	if err != nil {
		fmt.Println(sql)
		panic(fmt.Sprintf("selecting from reserve: %v", err))
	}
	defer rows.Close()
	result := ccms.NewResult("select")
	result.AddField("id", "bigint")
	result.AddField("author", "text")
	result.AddField("title", "text")
	result.AddField("full_vendor_name", "text")
	result.AddField("availability", "text")
	var count int
	for rows.Next() {
		var id int64
		var author, title, full_vendor_name, availability string
		err = rows.Scan(&id, &author, &title, &full_vendor_name, &availability)
		if err != nil {
			panic(fmt.Sprintf("reading from reserve: %v", err))
		}
		result.AddData([]any{id, author, title, full_vendor_name, availability})
		count++
		if count > 10000 {
			return nil, errors.New("result set too large")
		}
	}
	if err = rows.Err(); err != nil {
		panic(fmt.Sprintf("reading from reserve: %v", err))
	}
	return result, nil
}

func processQuery(s *svr, rqid int64, query *ast.QueryClause) error {
	if !s.cat.SetExists(query.From) {
		return errors.New("set \"" + query.From + "\" does not exist")
	}

	return nil
}
