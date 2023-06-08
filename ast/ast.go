package ast

type Node interface {
	GetType() string
}

type Stmt interface {
	Node
	GetType() string
}

type Expr interface {
	Stmt
	GetType() string
}

type BlockStmt struct {
	Stmt
	Body []Stmt
}

func (s BlockStmt) GetType() string {
	return "BlockStmt"
}

type Program struct {
	Stmt
	Body []Stmt
}

func (s Program) GetType() string {
	return "Program"
}