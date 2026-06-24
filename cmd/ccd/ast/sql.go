package ast

import (
	"errors"
	"fmt"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/cat"
	"github.com/indexdata/ccms/internal/dbx"
	"github.com/indexdata/ccms/internal/pgerr"
	"github.com/indexdata/ccms/internal/set"
	"github.com/jackc/pgx/v5"
)

// conversion to SQL

func (s *CreateFilterStmt) SQL(d *dbx.DB) (string, error) {
	var b strings.Builder
	if err := s.sql(d, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *CreateFilterStmt) sql(d *dbx.DB, b *strings.Builder) error {
	w := s.Where.(*WhereClause)
	if w.Valid {
		if err := evalExpr(d, b, w.Condition, true); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeleteStmt) SQL(d *dbx.DB) (string, error) {
	var b strings.Builder
	if err := s.sql(d, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *DeleteStmt) sql(d *dbx.DB, b *strings.Builder) error {
	fromSet := set.Parse(s.From)

	fromTable := cat.SetTable(fromSet)
	table := dbx.ParseTable(fromTable)

	b.WriteString("delete from ")
	b.WriteString(fromTable)
	w := s.Where.(*WhereClause)
	if w.Valid {
		b.WriteString(" where id in (")
		b.WriteString("select t.id from ")
		b.WriteString(cat.SetTable(fromSet))

		b.WriteString(" t join ccms.attr a on t.id=a.id")

		b.WriteString(" left join " + table.Schema + ".object o on t.id=o.id")
		b.WriteString(" left join ccms.fund on o.fund_id=fund.id")

		b.WriteString(" where (")
		if err := evalExpr(d, b, w.Condition, true); err != nil {
			return err
		}
		b.WriteRune(')')

		b.WriteRune(')')
	}
	return nil
}

func (s *InsertStmt) SQL(d *dbx.DB) (string, error) {
	var b strings.Builder
	if err := s.sql(d, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *InsertStmt) sql(d *dbx.DB, b *strings.Builder) error {
	intoSet := set.Parse(s.Into)

	b.WriteString("insert into ")
	b.WriteString(cat.SetTable(intoSet))
	b.WriteString(" select a.id ")
	if err := s.Query.(*QueryClause).sql(d, b); err != nil {
		return err
	}
	b.WriteString(" on conflict do nothing")
	return nil
}

func (s *SelectStmt) SQL(d *dbx.DB) (string, error) {
	var b strings.Builder
	if err := s.sql(d, &b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *SelectStmt) sql(d *dbx.DB, b *strings.Builder) error {
	var projection string
	switch s.AttrList.(*SelectAttrList).Attr {
	case "*":
		// projection = "a.id, coalesce(a.author, '') as author, coalesce(a.title, '') as title, coalesce(a.full_vendor_name, '') as full_vendor_name, coalesce(a.availability, '') as availability, coalesce(fund.name, '') fund"
		projection = "a.id, a.author, a.title, a.full_vendor_name, a.availability, o.decision, fund.name||':'||fund.title fund"
	case "count(*)":
		projection = "count(*)"
	}

	b.WriteString("select ")
	b.WriteString(projection)
	b.WriteRune(' ')
	if err := s.Query.(*QueryClause).sql(d, b); err != nil {
		return err
	}
	return nil
}

func (s *QueryClause) sql(d *dbx.DB, b *strings.Builder) error {
	fromSet := set.Parse(s.From)

	fromTable := cat.SetTable(fromSet)
	table := dbx.ParseTable(fromTable)

	b.WriteString("from ")
	if table.Table == "object" {
		b.WriteString("ccms.reserve")
	} else {
		b.WriteString(fromTable)
	}
	b.WriteString(" t join ccms.attr a on t.id=a.id")

	b.WriteString(" left join " + table.Schema + ".object o on t.id=o.id")
	b.WriteString(" left join ccms.fund on o.fund_id=fund.id")

	w := s.Where.(*WhereClause)
	if w.Valid {
		b.WriteString(" where (")
		if err := evalExpr(d, b, w.Condition, true); err != nil {
			return err
		}
		b.WriteRune(')')
	}
	o := s.Order.(*OrderClause)
	if o.Valid {
		b.WriteString(" order by ")
		b.WriteString(o.Attr)
		if o.Desc {
			b.WriteString(" desc")
		}
	}
	if s.Limit.(*LimitClause).Valid {
		b.WriteString(" limit ")
		b.WriteString(s.Limit.(*LimitClause).Count)
	}
	if s.Offset.(*OffsetClause).Valid {
		b.WriteString(" offset ")
		b.WriteString(s.Offset.(*OffsetClause).Start)
	}
	return nil
}

func (u *UpdateStmt) SQL() (string, error) {
	var b strings.Builder
	if err := u.sql(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (u *UpdateStmt) sql(b *strings.Builder) error {
	if u.ValueNull {
		b.WriteString("update ")
		b.WriteString(u.Set)
		b.WriteString(" set ")
		b.WriteString(u.Attr)
		b.WriteString("=null where id=")
		b.WriteString(u.IDValue.Value)
	} else {
		b.WriteString("insert into ")
		b.WriteString(u.Set)
		b.WriteString(" (id, ")
		b.WriteString(u.Attr)
		b.WriteString(") values (")
		b.WriteString(u.IDValue.Value)
		b.WriteString(", ")
		b.WriteString(u.Value)
		b.WriteString(")")
		b.WriteString(" on conflict (id) do update set ")
		b.WriteString(u.Attr)
		b.WriteRune('=')
		b.WriteString(u.Value)
	}
	return nil
}

func evalExpr(d *dbx.DB, b *strings.Builder, expr Node, root bool) error {
	switch e := expr.(type) {
	case *OrExpr:
		if err := evalExpr(d, b, e.Expr1, false); err != nil {
			return err
		}
		b.WriteString(" or ")
		if err := evalExpr(d, b, e.Expr2, false); err != nil {
			return err
		}
	case *AndExpr:
		if err := evalExpr(d, b, e.Expr1, false); err != nil {
			return err
		}
		b.WriteString(" and ")
		if err := evalExpr(d, b, e.Expr2, false); err != nil {
			return err
		}
	case *NotExpr:
		b.WriteString("not ")
		if err := evalExpr(d, b, e.Expr, false); err != nil {
			return err
		}
	case *EqualExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('=')
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *LikeExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" like ")
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *ILikeExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" ilike ")
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *NotEqualExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteString("<>")
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *LessThanExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('<')
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *GreaterThanExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('>')
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *LessThanOrEqualExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteString("<=")
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *GreaterThanOrEqualExpr:
		if err := evalExprOptAttr(d, b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(">=")
		if err := evalExprOptAttr(d, b, e.Expr2); err != nil {
			return err
		}
	case *FilterExpr:
		b.WriteRune('(')

		// b.WriteString(e.Filter)
		// if err := evalExprList(b, e.ExprList); err != nil {
		// 	return err
		// }

		rows, _ := d.Q.Query(d.C, "select sql from ccms.filter where name=$1", e.Filter)
		filter, err := pgx.CollectRows(rows, pgx.RowTo[string])
		if err != nil {
			return pgerr.Error(err)
		}
		b.WriteString(filter[0])

		b.WriteRune(')')
	case *TagExpr:
		return fmt.Errorf("tag() is not yet supported")
		//b.WriteString("TAG(")
		//if err := evalExprList(b, e.ExprList); err != nil {
		//        return err
		//}
		//b.WriteRune(')')
	case *Name:
		if root {
			return errors.New("invalid boolean expression")
		}
		b.WriteString(e.Value)
	case *SLiteral:
		if root {
			return errors.New("invalid boolean expression")
		}
		b.WriteRune('\'')
		b.WriteString(e.Value)
		b.WriteRune('\'')
	case *Number:
		if root {
			return errors.New("invalid boolean expression")
		}
		b.WriteString(e.Value)
	case *ParenExpr:
		b.WriteRune('(')
		if err := evalExpr(d, b, e.Expr, root); err != nil {
			return err
		}
		b.WriteRune(')')
	default:
		return fmt.Errorf("unknown node %T", expr)
	}
	return nil
}

func evalExprList(d *dbx.DB, b *strings.Builder, exprList []Node) error {
	for i := range exprList {
		if i != 0 {
			b.WriteRune(',')
		}
		if err := evalExpr(d, b, exprList[i], false); err != nil {
			return err
		}
	}
	return nil
}

// evaluate expr which may optionally be an attribute
// if expr is of type Name, require that it be a valid attribute name
func evalExprOptAttr(d *dbx.DB, b *strings.Builder, expr Node) error {
	switch e := expr.(type) {
	case *Name:
		if cat.IsAttribute(e.Value) {
			attrSQL(b, e.Value)
		} else {
			if e.Value == "true" || e.Value == "false" {
				b.WriteString(e.Value)
			} else {
				b.WriteRune('\'')
				b.WriteString(e.Value)
				b.WriteRune('\'')
				// return errors.New("attribute \"" + e.Value + "\" does not exist")
			}
		}
	default:
		if err := evalExpr(d, b, expr, false); err != nil {
			return err
		}
	}
	return nil
}

func attrSQL(b *strings.Builder, attr string) {
	switch attr {
	case "id", "author", "title", "full_vendor_name", "availability":
		b.WriteRune('a')
		b.WriteRune('.')
		b.WriteString(attr)
	case "decision":
		b.WriteString("coalesce(")
		b.WriteRune('o')
		b.WriteRune('.')
		b.WriteString(attr)
		b.WriteString(",false)")
	case "fund":
		b.WriteString("coalesce(fund.name,'')")
	default:
		b.WriteString(attr)
	}
}
