package ast

type Stmt interface {
	AstNode
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