package ast

type Node interface {
	GetType() string
}

type Stmt interface {
	Node
}

type Expr interface {
	Node
}