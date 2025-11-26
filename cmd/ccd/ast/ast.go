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

type ShowFiltersStmt struct {
}

func (*ShowFiltersStmt) node()     {}
func (*ShowFiltersStmt) stmtNode() {}

type ShowSetsStmt struct {
}

func (*ShowSetsStmt) node()     {}
func (*ShowSetsStmt) stmtNode() {}

type PingStmt struct {
}

func (*PingStmt) node()     {}
func (*PingStmt) stmtNode() {}
