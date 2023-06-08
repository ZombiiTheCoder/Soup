package ast

type VarDec struct {
	Stmt
	NotMut bool
	Name   string
	Val    any
}

func (s VarDec) GetType() string {
	return "VarDec"
}

type FuncDec struct {
	Stmt
	Name   string
	Params []string
	Body   BlockStmt
}

func (s FuncDec) GetType() string {
	return "FuncDec"
}

type ImpStmt struct {
	Stmt
	File string
	Rel  bool
}

func (s ImpStmt) GetType() string {
	return "ImpStmt"
}

type ForStmt struct {
	Stmt
	Init   any
	Test   any
	Update any
	Body   BlockStmt
}

func (s ForStmt) GetType() string {
	return "ForStmt"
}

type IfStmt struct {
	Stmt
	Test      Expr
	Consquent BlockStmt
	Alternate any
}

func (s IfStmt) GetType() string {
	return "IfStmt"
}

type WhileStmt struct {
	Stmt
	Test      Expr
	Consquent BlockStmt
}

func (s WhileStmt) GetType() string {
	return "WhileStmt"
}

type ReturnStmt struct {
	Stmt
	Value Expr
}

func (s ReturnStmt) GetType() string {
	return "ReturnStmt"
}