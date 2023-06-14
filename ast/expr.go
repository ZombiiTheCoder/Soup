package ast

type AssignExpr struct {
	Expr
	Type     string
	Assigner Expr
	Val      Expr
	Op       string
}

type TernaryExpr struct {
	Expr
	Type      string
	Condition Expr
	Consquent Expr
	Alternate any
}

type LogicalExpr struct {
	Expr
	Type  string
	Op    string
	Left  Expr
	Right Expr
}

type BinaryExpr struct {
	Expr
	Type  string
	Op    string
	Left  Expr
	Right Expr
}

type UnaryExpr struct {
	Expr
	Type     string
	Op       string
	Prefix   bool
	Argument Expr
}

type CallExpr struct {
	Expr
	Type   string
	Caller Expr
	Args   []Expr
}

type MemberExpr struct {
	Expr
	Type     string
	Computed bool
	Property Expr
	Obj      Expr
}

type ArrayExpr struct {
	Expr
	Type     string
	Elements []Expr
}

type Property struct {
	Expr
	Type string
	Key  string
	Val  any
}