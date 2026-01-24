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
	Select   SelectExpr
	Set      string
	Limit    string
	Retrieve bool
}

func (*SelectStmt) node()     {}
func (*SelectStmt) stmtNode() {}

type ShowStmt struct {
	Name string
}

func (*ShowStmt) node()     {}
func (*ShowStmt) stmtNode() {}

type PingStmt struct {
}

func (*PingStmt) node()     {}
func (*PingStmt) stmtNode() {}
