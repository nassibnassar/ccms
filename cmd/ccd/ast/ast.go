package ast

import "github.com/indexdata/ccms/cmd/ccd/catalog"

type Option struct {
	Action string
	Name   string
	Val    string
}

type Node interface {
	node()
}

type Expr interface {
	Node
	exprNode()
}

type WhereExpr interface {
	Node
	exprNode()
	whereExprNode()
}

type NoWhereExpr struct {
}

func (*NoWhereExpr) node()          {}
func (*NoWhereExpr) exprNode()      {}
func (*NoWhereExpr) whereExprNode() {}

type WhereConditionExpr struct {
	WhereAttr  string
	WhereValue string
}

func (*WhereConditionExpr) node()          {}
func (*WhereConditionExpr) exprNode()      {}
func (*WhereConditionExpr) whereExprNode() {}

type OrderExpr interface {
	Node
	exprNode()
	orderExprNode()
}

type NoOrderExpr struct {
}

func (*NoOrderExpr) node()          {}
func (*NoOrderExpr) exprNode()      {}
func (*NoOrderExpr) orderExprNode() {}

type OrderValueExpr struct {
	Attribute string
	Desc      bool
}

func (*OrderValueExpr) node()          {}
func (*OrderValueExpr) exprNode()      {}
func (*OrderValueExpr) orderExprNode() {}

type QueryExpr struct {
	From  string
	Where WhereExpr
	Order OrderExpr
	Limit LimitExpr
}

func (*QueryExpr) node()     {}
func (*QueryExpr) exprNode() {}

type SelectExpr interface {
	Node
	exprNode()
	selectExprNode()
}

type StarSelectExpr struct {
}

func (*StarSelectExpr) node()           {}
func (*StarSelectExpr) exprNode()       {}
func (*StarSelectExpr) selectExprNode() {}

type AttrSelectExpr struct {
	Attribute string
}

func (*AttrSelectExpr) node()           {}
func (*AttrSelectExpr) exprNode()       {}
func (*AttrSelectExpr) selectExprNode() {}

type LimitExpr interface {
	Node
	exprNode()
	limitExprNode()
}

type NoLimitExpr struct {
}

func (*NoLimitExpr) node()          {}
func (*NoLimitExpr) exprNode()      {}
func (*NoLimitExpr) limitExprNode() {}

type LimitValueExpr struct {
	Value string
}

func (*LimitValueExpr) node()          {}
func (*LimitValueExpr) exprNode()      {}
func (*LimitValueExpr) limitExprNode() {}

type Stmt interface {
	Node
	stmtNode()
}

type InfoStmt struct {
	Topic string
}

func (*InfoStmt) node()     {}
func (*InfoStmt) stmtNode() {}

type CreateSetStmt struct {
	SetName string
}

func (*CreateSetStmt) node()     {}
func (*CreateSetStmt) stmtNode() {}

type InsertStmt struct {
	Into  string
	Query *QueryExpr
}

func (*InsertStmt) node()     {}
func (*InsertStmt) stmtNode() {}

type SelectStmt struct {
	Select SelectExpr
	Query  *QueryExpr
}

func (*SelectStmt) node()     {}
func (*SelectStmt) stmtNode() {}

func (q *QueryExpr) SQL() string {
	var where string
	switch w := q.Where.(type) {
	case *NoWhereExpr:
	case *WhereConditionExpr:
		where = " and " + w.WhereAttr + "='" + w.WhereValue + "'"
	}
	var order string
	switch o := q.Order.(type) {
	case *NoOrderExpr:
	case *OrderValueExpr:
		if o.Desc {
			order = " order by " + o.Attribute + " desc"
		} else {
			order = " order by " + o.Attribute
		}
	}
	var limit string
	switch l := q.Limit.(type) {
	case *NoLimitExpr:
	case *LimitValueExpr:
		limit = " limit " + l.Value
	}
	return "from " + catalog.SetTable(q.From) + " t join ccms.attr a on t.id=a.id" + where + order + limit
}

func (i *InsertStmt) SQL() string {
	return "insert into " + i.Into + " select a.id " + i.Query.SQL() + " on conflict do nothing"
}

func (s *SelectStmt) SQL() string {
	//var sel []string
	//switch se := cmd.Select.(type) {
	//case *AttrSelectExpr:
	//        sel = []string{se.Attribute}
	//case *StarSelectExpr:
	//        sel = []string{"*"}
	//}
	return "select a.id, coalesce(a.author, '') as author, coalesce(a.title, '') as title, coalesce(a.full_vendor_name, '') as full_vendor_name, coalesce(a.availability, '') as availability " + s.Query.SQL()
}

type ShowStmt struct {
	Name string
}

func (*ShowStmt) node()     {}
func (*ShowStmt) stmtNode() {}

type PingStmt struct {
}

func (*PingStmt) node()     {}
func (*PingStmt) stmtNode() {}
