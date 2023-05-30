package kinds

type AstNode interface {
	AsString() string
}

// Statements
type Program AstNode
type VarDec AstNode
type FuncDec AstNode
type IfState AstNode
type WhileState AstNode
type BlockState AstNode

// Expressions
type CallExpr AstNode
type MemberExpr AstNode
type UnaryExpr AstNode
type BinaryExpr AstNode
type AssignExpr AstNode

// Literals
type Property AstNode
type Identifier AstNode
type String AstNode
type ObjectLiteral AstNode
type NumericLiteral AstNode
type ListLiteral AstNode

// type AstNode int

// // Statements
// const (
// 	Program AstNode = iota
// 	VarDec
// 	FuncDec
// 	IfState
// 	WhileState
// 	BlockState

// // Expressions
// 	CallExpr
// 	MemberExpr
// 	UnaryExpr
// 	BinaryExpr
// 	AssignExpr

// // Literals
// 	Property
// 	Identifier
// 	String
// 	ObjectLiteral
// 	NumericLiteral
// )
