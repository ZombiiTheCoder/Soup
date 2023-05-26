package ast

type Stmt interface {
	AstNode
}

type BlockStmt struct {
	Stmt
	Body []Stmt
}

type Program struct {
	Stmt
	Body []Stmt
}

type VarDec struct {
	Stmt
	Cont bool
	Idnt string
	Valu any
}

type ImportDec struct {
	Stmt
	File string
}

type IfStmt struct{
	Stmt
	Condition Expr
	Consquent BlockStmt
	Alternate any
}

type WhileStmt struct{
	Stmt
	Condition Expr
	Consquent BlockStmt
}