package ast

type AstNode interface{
	GetType() string
	GetValue() string
}

type null *struct{}
var nul null