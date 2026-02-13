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
	b.WriteString(i.Into)
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
	b.WriteString("select a.id, coalesce(a.author, '') as author, coalesce(a.title, '') as title, coalesce(a.full_vendor_name, '') as full_vendor_name, coalesce(a.availability, '') as availability ")
	if err := s.Query.(*QueryClause).sql(b); err != nil {
		return err
	}
	return nil
}

func (q *QueryClause) sql(b *strings.Builder) error {
	b.WriteString("from ")
	b.WriteString(catalog.SetTable(q.From))
	b.WriteString(" t join ccms.attr a on t.id=a.id")
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
