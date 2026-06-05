package ast

import (
	"errors"
	"fmt"
	"strings"

	"github.com/indexdata/ccms/cmd/ccd/catalog"
)

// conversion to SQL

func (d *DeleteStmt) SQL() (string, error) {
	var b strings.Builder
	if err := d.sql(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (d *DeleteStmt) sql(b *strings.Builder) error {
	b.WriteString("delete from ")
	b.WriteString(catalog.SetTable(d.From))
	w := d.Where.(*WhereClause)
	if w.Valid {
		b.WriteString(" t using ccms.attr a where t.id=a.id and (")
		if err := evalExpr(b, w.Condition); err != nil {
			return err
		}
		b.WriteRune(')')
	}
	return nil
}

func (i *InsertStmt) SQL() (string, error) {
	var b strings.Builder
	if err := i.sql(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (i *InsertStmt) sql(b *strings.Builder) error {
	b.WriteString("insert into ")
	b.WriteString(catalog.SetTable(i.Into))
	b.WriteString(" select a.id ")
	if err := i.Query.(*QueryClause).sql(b); err != nil {
		return err
	}
	b.WriteString(" on conflict do nothing")
	return nil
}

func (s *SelectStmt) SQL() (string, error) {
	var b strings.Builder
	if err := s.sql(&b); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (s *SelectStmt) sql(b *strings.Builder) error {
	var projection string
	switch s.AttrList.(*SelectAttrList).Attr {
	case "*":
		// projection = "a.id, coalesce(a.author, '') as author, coalesce(a.title, '') as title, coalesce(a.full_vendor_name, '') as full_vendor_name, coalesce(a.availability, '') as availability, coalesce(fund.name, '') fund"
		projection = "a.id, a.author, a.title, a.full_vendor_name, a.availability, fund.name||':'||fund.title fund"
	case "count(*)":
		projection = "count(*)"
	}

	b.WriteString("select ")
	b.WriteString(projection)
	b.WriteRune(' ')
	if err := s.Query.(*QueryClause).sql(b); err != nil {
		return err
	}
	return nil
}

func (q *QueryClause) sql(b *strings.Builder) error {
	fromTable := catalog.SetTable(q.From)
	schema, table := catalog.SplitSchemaTable(fromTable)

	b.WriteString("from ")
	if table == "object" {
		b.WriteString("ccms.reserve")
	} else {
		b.WriteString(fromTable)
	}
	b.WriteString(" t join ccms.attr a on t.id=a.id")

	b.WriteString(" left join " + schema + ".object on t.id=object.id")
	b.WriteString(" left join ccms.fund on object.fund_id=fund.id")

	w := q.Where.(*WhereClause)
	if w.Valid {
		b.WriteString(" where (")
		if err := evalExpr(b, w.Condition); err != nil {
			return err
		}
		b.WriteRune(')')
	}
	o := q.Order.(*OrderClause)
	if o.Valid {
		b.WriteString(" order by ")
		b.WriteString(o.Attr)
		if o.Desc {
			b.WriteString(" desc")
		}
	}
	if q.Limit.(*LimitClause).Valid {
		b.WriteString(" limit ")
		b.WriteString(q.Limit.(*LimitClause).Count)
	}
	if q.Offset.(*OffsetClause).Valid {
		b.WriteString(" offset ")
		b.WriteString(q.Offset.(*OffsetClause).Start)
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
	if u.Value == "" {
		b.WriteString("update ")
		b.WriteString(u.SetName)
		b.WriteString(" set fund_id=null where id=")
		b.WriteString(u.IDValue.Value)
	} else {
		b.WriteString("insert into ")
		b.WriteString(u.SetName)
		b.WriteString(" (id, fund_id) values (")
		b.WriteString(u.IDValue.Value)
		b.WriteString(", ")
		b.WriteString(u.Value)
		b.WriteString(")")
		b.WriteString(" on conflict (id) do update set fund_id=")
		b.WriteString(u.Value)
	}
	return nil
}

func evalExpr(b *strings.Builder, expr Node) error {
	switch e := expr.(type) {
	case *OrExpr:
		if err := evalExpr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" or ")
		if err := evalExpr(b, e.Expr2); err != nil {
			return err
		}
	case *AndExpr:
		if err := evalExpr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" and ")
		if err := evalExpr(b, e.Expr2); err != nil {
			return err
		}
	case *NotExpr:
		b.WriteString("not ")
		if err := evalExpr(b, e.Expr); err != nil {
			return err
		}
	case *EqualExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('=')
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *LikeExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" like ")
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *ILikeExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(" ilike ")
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *NotEqualExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString("<>")
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *LessThanExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('<')
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *GreaterThanExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteRune('>')
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *LessThanOrEqualExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString("<=")
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *GreaterThanOrEqualExpr:
		if err := evalExprOptAttr(b, e.Expr1); err != nil {
			return err
		}
		b.WriteString(">=")
		if err := evalExprOptAttr(b, e.Expr2); err != nil {
			return err
		}
	case *FilterExpr:
		return fmt.Errorf("filter() is not yet supported")
		//b.WriteString("FILTER(")
		//if err := evalExprList(b, e.ExprList); err != nil {
		//        return err
		//}
		//b.WriteRune(')')
	case *TagExpr:
		return fmt.Errorf("tag() is not yet supported")
		//b.WriteString("TAG(")
		//if err := evalExprList(b, e.ExprList); err != nil {
		//        return err
		//}
		//b.WriteRune(')')
	case *Name:
		b.WriteString(e.Value)
	case *SLiteral:
		b.WriteRune('\'')
		b.WriteString(e.Value)
		b.WriteRune('\'')
	case *Number:
		b.WriteString(e.Value)
	case *ParenExpr:
		b.WriteRune('(')
		if err := evalExpr(b, e.Expr); err != nil {
			return err
		}
		b.WriteRune(')')
	default:
		return fmt.Errorf("internal error: unknown node %T", expr)
	}
	return nil
}

func evalExprList(b *strings.Builder, exprList []Node) error {
	for i := range exprList {
		if i != 0 {
			b.WriteRune(',')
		}
		if err := evalExpr(b, exprList[i]); err != nil {
			return err
		}
	}
	return nil
}

// evaluate expr which may optionally be an attribute
// if expr is of type Name, require that it be a valid attribute name
func evalExprOptAttr(b *strings.Builder, expr Node) error {
	switch e := expr.(type) {
	case *Name:
		if !catalog.IsAttr(e.Value) {
			return errors.New("attribute \"" + e.Value + "\" does not exist")
		}
		b.WriteString("a.")
	}
	if err := evalExpr(b, expr); err != nil {
		return err
	}
	return nil
}
