package ast

type Identifier struct {
	Expr
	Type string `json:Type`
	Symb string `json:Symb`
}

type StringLiteral struct {
	Expr
	Type string `json:Type`
	Valu string `json:Valu`
}

type IntegerLiteral struct {
	Expr
	Type string `json:Type`
	Valu int64  `json:Valu`
}

type FloatLiteral struct {
	Expr
	Type string  `json:Type`
	Valu float64 `json:Valu`
}

type NullLiteral struct {
	Expr
	Type string `json:Type`
	Valu string `json:Valu`
}

type ObjectLiteral struct {
	Expr
	Type       string     `json:Type`
	Properties []Property `json:Properties`
}