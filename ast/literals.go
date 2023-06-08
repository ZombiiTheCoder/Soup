package ast

type Identifier struct {
	Expr
	Symb string
}

func (s Identifier) GetType() string {
	return "Identifier"
}

type StringLiteral struct {
	Expr
	Valu string
}

func (s StringLiteral) GetType() string {
	return "String"
}

type IntegerLiteral struct {
	Expr
	Valu int64
}

func (s IntegerLiteral) GetType() string {
	return "Integer"
}

type FloatLiteral struct {
	Expr
	Valu float64
}

func (s FloatLiteral) GetType() string {
	return "Float"
}

type NullLiteral struct {
	Expr
	Valu string
}

func (s NullLiteral) GetType() string {
	return "Null"
}

type ObjectLiteral struct {
	Expr
	Properties []Property
}

func (s ObjectLiteral) GetType() string {
	return "Object"
}