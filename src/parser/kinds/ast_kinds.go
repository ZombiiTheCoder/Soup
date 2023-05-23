package kinds

type NodeType int

// Statements
type Program NodeType
type VarDec NodeType
type FuncDec NodeType
type IfState NodeType
type WhileState NodeType
type BlockState NodeType

// Expressions
type CallExpr NodeType
type MemberExpr NodeType
type UnaryExpr NodeType
type BinaryExpr NodeType
type AssignExpr NodeType

// Literals
type Property NodeType
type Identifier NodeType
type String NodeType
type ObjectLiteral NodeType
type NumericLiteral NodeType