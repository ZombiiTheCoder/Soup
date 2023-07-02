package ast

type TType struct {
	Expr
	NodeType string

	Value string
}

type Identifier struct {
	Expr
	NodeType string

	Value string
}