package parser

import (
	"strconv"
)

type AstNode interface{
	AsString() string
	GetType() string
}

type Stmnt interface {
	AstNode
}

type Expr interface {
	Stmnt
}

type Program struct {
	Stmnt
	Body []Stmnt
}

type VarDec struct {
	Stmnt
	Cont bool
	Idnt string
	Value any
}

func CreateVarDec (cont bool, ident string, value any) Stmnt {
	return VarDec{ Cont: cont, Idnt: ident, Value: value }
}

type AssignExpr struct {
	Expr
	Assigner Expr
	Valu Expr
}

func CreateAssignExpr (assigner, value Expr) Expr {
	return AssignExpr{ Assigner: assigner, Valu: value }
}

type Property struct {
	Expr
	Key string
	Val any
}

func CreateProperty (key string, value any) Property {
	return Property{ Key: key, Val: value }
}

type ObjectLiteral struct {
	Expr
	Properties []Property
}

func CreateObjectLiteral (properties []Property) Expr {
	return ObjectLiteral{ Properties: properties }
}

type MemberExpr struct {
	Expr
	Obj Expr
	Property Expr
	Computed bool
}

func CreateMemberExpr (object, property Expr, computed bool) Expr {
	return MemberExpr{ Obj: object, Property: property, Computed: computed }
}

type BinaryExpr struct {
	Expr
	Left Expr
	Right Expr
	Op string
}

func CreateBinaryExpr (left, right Expr, op string) Expr {
	return BinaryExpr{ Left: left, Right: right, Op: op }
}

type UnaryExpr struct {
	Expr
	Left Expr
	Op string
	Prefix bool
}

func CreateUnaryExpr (left Expr, op string, prefix bool) Expr {
	return UnaryExpr{ Left: left, Op: op, Prefix: prefix }
}

type Identifier struct {
	Expr
	Symb string
}

func CreateIdentifier (symbol string) Expr {
	return Identifier{ Symb: symbol }
}

type NumericLiteral struct {
	Expr
	Valu float64
}

func CreateNumericLiteral (value string) Expr {
	v, _:=strconv.ParseFloat(value, 64)
	return NumericLiteral{ Valu: v }
}

type StringLiteral struct {
	Expr
	Valu string
}

func CreateStringLiteral (value string) Expr {
	return StringLiteral{ Valu: value }
}

type CallExpr struct {
	Expr
	Caller Expr
	Args []Expr
}

func CreateCallExpr (caller Expr, args []Expr) Expr {
	return CallExpr{ Caller: caller, Args: args }
}