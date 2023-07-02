package ast

type Program struct {
	Stmt
	NodeType string

	Body []Stmt
}

type BlockStmt struct {
	Stmt
	NodeType string

	Type string
	Body []Stmt
}

type VarDeclaration struct {
	Stmt
	NodeType string

	Type  string
	Const bool
	Name  string
	Value any
}

type FuncDeclaration struct {
	Stmt
	NodeType string

	Type   string
	Name   string
	Params map[string]string
	Body   []Stmt
}

type IfStmt struct {
	Stmt
	NodeType string

	Type       string
	Condition  Stmt
	Consequent []Stmt
	Alternate  any
}

type WhileStmt struct {
	Stmt
	NodeType string

	Type       string
	Condition  Stmt
	Consequent []Stmt
}

type FuncDec struct {
	Stmt
	NodeType string

	Type   string
	Name   string
	Params map[string]string
	Body   []Stmt
}

type ReturnStmt struct {
	Stmt
	NodeType string

	Type  string
	Value Expr
}