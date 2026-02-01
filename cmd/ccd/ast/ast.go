package ast

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

type SelectStmt struct {
	Select     SelectExpr
	From       string
	WhereAttr  string
	WhereValue string
	Limit      LimitExpr
	Retrieve   bool
}

func (*SelectStmt) node()     {}
func (*SelectStmt) stmtNode() {}

func (s *SelectStmt) SQL() string {
	//var sel []string
	//switch se := cmd.Select.(type) {
	//case *AttrSelectExpr:
	//        sel = []string{se.Attribute}
	//case *StarSelectExpr:
	//        sel = []string{"*"}
	//}
	var where string
	if s.WhereAttr != "" {
		where = " and " + s.WhereAttr + "='" + s.WhereValue + "'"
	}
	var limit string
	switch l := s.Limit.(type) {
	case *NoLimitExpr:
	case *LimitValueExpr:
		limit = " limit " + l.Value
	}
	return "select a.id, coalesce(a.author, ''), coalesce(a.title, ''), coalesce(a.full_vendor_name, ''), coalesce(a.availability, '') from " + s.From + " t join ccms.attr a on t.id=a.id where a.author is not null" + where + limit
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
