package parser

import (
	"fmt"
	"strconv"
)

type null *struct{}
var niz null

type AstNode interface{
	GetType() string
	GetValue() string
}

type Stmt interface {
	AstNode
}

type Expr interface {
	Stmt
}

type Program struct {
	Stmt
	Body []Stmt
}

func (s Program) GetType() string {
	return "Program"
}

type VarDec struct {
	Stmt
	Cont bool
	Idnt string
	Valu any
}

func (s VarDec) GetType() string {
	return "VarDec"
}

func CreateVarDec (cont bool, ident string, value any) Stmt {
	return VarDec{ Cont: cont, Idnt: ident, Valu: value }
}

type AssignExpr struct {
	Expr
	Assigner Expr
	Valu Expr
}

func (s AssignExpr) GetType() string {
	return "AssignExpr"
}

func CreateAssignExpr (assigner, value Expr) Expr {
	return AssignExpr{ Assigner: assigner, Valu: value}
}

type Property struct {
	Expr
	Key string
	Val any
}

func (s Property) GetType() string {
	return "Property"
}

func CreateProperty (key string, value any) Property {
	return Property{ Key: key, Val: value }
}

type ObjectLiteral struct {
	Expr
	Properties []Property
}

func (s ObjectLiteral) GetType() string {
	return "ObjectLiteral"
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

func (s MemberExpr) GetType() string {
	return "MemberExpr"
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

func (s BinaryExpr) GetType() string {
	return "BinaryExpr"
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

func (s UnaryExpr) GetType() string {
	return "UnaryExpr"
}

func CreateUnaryExpr (left Expr, op string, prefix bool) Expr {
	return UnaryExpr{ Left: left, Op: op, Prefix: prefix }
}

type Identifier struct {
	Expr
	Symb string
}

func (s Identifier) GetType() string {
	return "Identifier"
}

func (s Identifier) GetValue() string{
	return s.Symb
}

func CreateIdentifier (symbol string) Expr {
	return Identifier{ Symb: symbol }
}

type NumericLiteral struct {
	Expr
	Valu float64
}

func (s NumericLiteral) GetType() string {
	return "NumericLiteral"
}

func (s NumericLiteral) GetValue() string {
	return string(fmt.Sprintf("%f", s.Valu))
}

func CreateNumericLiteral (value string) Expr {
	v, _:=strconv.ParseFloat(value, 64)
	return NumericLiteral{ Valu: v }
}

type StringLiteral struct {
	Expr
	Valu string
}

func (s StringLiteral) GetType() string {
	return "StringLiteral"
}

func (s StringLiteral) GetValue() string {
	return s.Valu
}

func CreateStringLiteral (value string) Expr {
	return StringLiteral{ Valu: value }
}

type NullLiteral struct {
	Expr
	Valu null
}

func (s NullLiteral) GetType() string {
	return "NullLiteral"
}

type CallExpr struct {
	Expr
	Caller Expr
	Args []Expr
}

func (s CallExpr) GetType() string {
	return "CallExpr"
}

func CreateCallExpr (caller Expr, args []Expr) Expr {
	return CallExpr{ Caller: caller, Args: args }
}