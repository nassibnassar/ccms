package ast

type Option struct {
	Action string
	Name   string
	Val    string
}

type Node interface {
	node()
}

type Stmt interface {
	Node
	stmtNode()
}

type HelpStmt struct {
}

func (*HelpStmt) node()     {}
func (*HelpStmt) stmtNode() {}

type CreateSetStmt struct {
	SetName string
}

func (*CreateSetStmt) node()     {}
func (*CreateSetStmt) stmtNode() {}

type ShowStmt struct {
	Name string
}

func (*ShowStmt) node()     {}
func (*ShowStmt) stmtNode() {}

type PingStmt struct {
}

func (*PingStmt) node()     {}
func (*PingStmt) stmtNode() {}
