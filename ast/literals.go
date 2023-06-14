package ast

type Identifier struct {
	Expr
	Type string
	Symb string
}

type StringLiteral struct {
	Expr
	Type string
	Valu string
}

type IntegerLiteral struct {
	Expr
	Type string
	Valu int64
}

type FloatLiteral struct {
	Expr
	Type string
	Valu float64
}

type NullLiteral struct {
	Expr
	Type string
	Valu string
}

type ObjectLiteral struct {
	Expr
	Type       string
	Properties []Property
}