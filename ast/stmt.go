package ast

type VarDec struct {
	Stmt
	Type   string `json:Type`
	NotMut bool   `json:NotMut`
	Name   string `json:Name`
	Value  any    `json:Value`
}

type FuncDec struct {
	Stmt
	Type   string   `json:Type`
	Name   string   `json:Name`
	Params []string `json:Params`
	Body   []Stmt   `json:Body`
}

type ImpStmt struct {
	Stmt
	Type string `json:Type`
	File string `json:File`
	Rel  bool   `json:Rel`
}

type IfStmt struct {
	Stmt
	Type      string `json:Type`
	Test      Expr   `json:Test`
	Consquent []Stmt `json:Consquent`
	Alternate any    `json:Alternate`
}

type WhileStmt struct {
	Stmt
	Type      string `json:Type`
	Test      Expr   `json:Test`
	Consquent []Stmt `json:Consquent`
}

type ReturnStmt struct {
	Stmt
	Type  string `json:Type`
	Value Expr   `json:Value`
}