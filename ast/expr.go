package ast

type AssignExpr struct {
	Expr
	Type     string `json:Type`
	Assigner Expr
	Val      Expr
	Op       string
}

type TernaryExpr struct {
	Expr
	Type      string `json:Type`
	Condition Expr
	Consquent Expr
	Alternate any
}

type LogicalExpr struct {
	Expr
	Type  string `json:Type`
	Op    string
	Left  Expr
	Right Expr
}

type BinaryExpr struct {
	Expr
	Type  string `json:Type`
	Op    string `json:Op`
	Left  Expr   `json:Left`
	Right Expr   `json:Right`
}

type UnaryExpr struct {
	Expr
	Type     string `json:Type`
	Op       string `json:Op`
	Prefix   bool   `json:Prefix`
	Argument Expr   `json:Argument`
}

type CallExpr struct {
	Expr
	Type   string `json:Type`
	Caller Expr   `json:Caller`
	Args   []Expr `json:Args`
}

type MemberExpr struct {
	Expr
	Type     string `json:Type`
	Computed bool   `json:Computed`
	Property Expr   `json:Property`
	Obj      Expr   `json:Obj`
}

type ArrayExpr struct {
	Expr
	Type     string `json:Type`
	Elements []Expr `json:Elements`
}

type Property struct {
	Expr
	Type string `json:Type`
	Key  string `json:Key`
	Val  any    `json:Val`
}