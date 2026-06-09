package server

import (
	"errors"
	"strconv"

	"github.com/indexdata/ccms"
	"github.com/indexdata/ccms/cmd/ccd/ast"
	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/internal/set"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
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

	from := cmd.Query.(*ast.QueryClause).From
	if from == "reserve" { // TODO remove this "reserve" check after some time
		return cmderr("set \"reserve\" is no longer supported; use \"<project>.object\"")
	}
	fromSet := set.Parse(from)
	setExists, err := cat.SetExists(s.d, fromSet)
	if err != nil {
		return cmderrint("checking if set exists", err)
	}
	if !setExists {
		return cmderr("set \"" + from + "\" does not exist")
	}

	o := cmd.Query.(*ast.QueryClause).Order.(*ast.OrderClause)
	if o.Valid {
		if !cat.IsAttr(o.Attr) {
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
		return cmderr(internalError + "invalid projection in select")
	}
}

func runQueryCount(s *svr, sql string) (*ccms.Result, error) {
	var count int64
	if err := s.d.Q.QueryRow(s.d.C, sql).Scan(&count); err != nil {
		return nil, errors.New(internalError + pgerr.Error(err).Error())
	}
	result := ccms.NewResult("select")
	result.AddField("count", "bigint")
	result.AddData([]any{count})
	return result, nil
}

func runQuery(s *svr, sql string) (*ccms.Result, error) {
	rows, err := s.d.Q.Query(s.d.C, sql)
	if err != nil {
		return nil, errors.New(internalError + pgerr.Error(err).Error())
	}
	defer rows.Close()
	result := ccms.NewResult("select")
	result.AddField("id", "bigint")
	result.AddField("author", "text")
	result.AddField("title", "text")
	result.AddField("full_vendor_name", "text")
	result.AddField("availability", "text")
	result.AddField("fund", "text")
	var count int
	for rows.Next() {
		var id int64
		// var author, title, full_vendor_name, availability, fund string
		var author, title, full_vendor_name, availability, fund zeronull.Text
		err = rows.Scan(&id, &author, &title, &full_vendor_name, &availability, &fund)
		if err != nil {
			return nil, errors.New(internalError + pgerr.Error(err).Error())
		}
		result.AddData([]any{id, author, title, full_vendor_name, availability, fund})
		count++
		if count > 10000 {
			return nil, errors.New("result set too large")
		}
	}
	if err = rows.Err(); err != nil {
		return nil, errors.New(internalError + pgerr.Error(err).Error())
	}
	return result, nil
}
