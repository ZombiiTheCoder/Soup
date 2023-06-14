package ast

type VarDec struct {
	Stmt
	Type   string
	NotMut bool
	Name   string
	Value  any
}

type FuncDec struct {
	Stmt
	Type   string
	Name   string
	Params []string
	Body   []Stmt
}

type ImpStmt struct {
	Stmt
	Type string
	File string
	Rel  bool
}

type IfStmt struct {
	Stmt
	Type      string
	Test      Expr
	Consquent []Stmt
	Alternate any
}

type WhileStmt struct {
	Stmt
	Type      string
	Test      Expr
	Consquent []Stmt
}

type ReturnStmt struct {
	Stmt
	Type  string
	Value Expr
}