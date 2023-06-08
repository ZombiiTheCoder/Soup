package ast

type AssignExpr struct {
	Expr
	Assigner Expr
	Val      Expr
	Op       string
}

func (s AssignExpr) GetType() string {
	return "AssignExpr"
}

type TernaryExpr struct {
	Expr
	Condition Expr
	Consquent Expr
	Alternate any
}

func (s TernaryExpr) GetType() string {
	return "TernaryExpr"
}

type LogicalExpr struct {
	Expr
	Op    string
	Left  Expr
	Right Expr
}

func (s LogicalExpr) GetType() string {
	return "LogicalExpr"
}

type BinaryExpr struct {
	Expr
	Op    string
	Left  Expr
	Right Expr
}

func (s BinaryExpr) GetType() string {
	return "BinaryExpr"
}

type UnaryExpr struct {
	Expr
	Op       string
	Prefix   bool
	Argument Expr
}

func (s UnaryExpr) GetType() string {
	return "UnaryExpr"
}

type CallExpr struct {
	Expr
	Caller Expr
	Args   []Expr
}

func (s CallExpr) GetType() string {
	return "CallExpr"
}

type MemberExpr struct {
	Expr
	Computed bool
	Property Expr
	Obj      Expr
}

func (s MemberExpr) GetType() string {
	return "MemberExpr"
}

type ArrayExpr struct {
	Expr
	Elements []Expr
}

func (s ArrayExpr) GetType() string {
	return "ArrayExpr"
}

type Property struct {
	Expr
	Key string
	Val any
}