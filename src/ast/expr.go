package ast

type Expr interface {
	Stmt
}

type AssignExpr struct {
	Expr
	Assigner Expr
	Valu     Expr
	Op       string
}

type Property struct {
	Expr
	Key string
	Val any
}

type ObjectLiteral struct {
	Expr
	Properties []Property
}

type MemberExpr struct {
	Expr
	Obj      Expr
	Property Expr
	Computed bool
}

type BinaryExpr struct {
	Expr
	Left  Expr
	Right Expr
	Op    string
}

type UnaryExpr struct {
	Expr
	Left   Expr
	Op     string
	Prefix bool
	Var    string
}

type CallExpr struct {
	Expr
	Caller Expr
	Args   []Expr
}

type Identifier struct {
	Expr
	Symb string
}

type StringLiteral struct {
	Expr
	Valu string
}

type NumericLiteral struct {
	Expr
	Valu int
}

type FloatLiteral struct {
	Expr
	Valu float64
}

type NullLiteral struct {
	Expr
	Valu null
}

type RelationalExpr struct {
	Expr
	Left  Expr
	Right Expr
	Op    string
}

type ArrayExpr struct {
	Expr
	Elements []Expr
}