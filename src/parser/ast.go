package parser

import "Soup/src/parser/kinds"

type Stmnt struct {
	kind kinds.NodeType
}

type Expr struct {
	Stmnt
}

type Program struct {
	Stmnt
	kind kinds.Program
	body []Stmnt
}

type VarDec struct {
	Stmnt
	kind kinds.VarDec
	cont bool
	Idnt string
	valu any
}

type AssignExpr struct {
	Expr
	kind kinds.AssignExpr
	assigner Expr
	valu Expr
}

type Property struct {
	Expr
	kind kinds.Property
	key string
	valu any
}