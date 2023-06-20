package ast

type Node interface {
	GetType() string
}

type Stmt interface {
	Node
}

type Expr interface {
	Stmt
}

type BlockStmt struct {
	Stmt
	Type string `json:Type`
	Body []Stmt `json:Body`
}

type Program struct {
	Stmt
	Type string `json:Type`
	Body []Stmt `json:Body`
}