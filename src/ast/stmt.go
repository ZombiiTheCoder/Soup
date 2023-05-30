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
	Cont    bool
	Idnt    string
	Valu    any
	PkgName string
}

type FuncDec struct {
	Stmt
	Name    string
	Params  []string
	Body    []Stmt
	PkgName string
}

type RetStmt struct {
	Stmt
	Valu Expr
}

type ImpStmt struct {
	Stmt
	File string
	Rel  bool
}

type IfStmt struct {
	Stmt
	Condition Expr
	Consquent BlockStmt
	Alternate any
}

type WhileStmt struct {
	Stmt
	Condition Expr
	Consquent BlockStmt
}
