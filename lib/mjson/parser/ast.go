package parser

type Node interface {
	GetType() string
}

type Stmt interface {
	Node
}

type Program struct {
	Stmt
	Type string
	Body Stmt
}

type Property struct {
	Stmt
	Type  string
	Key   string
	Value Stmt
}

type Object struct {
	Stmt
	Type       string
	Properties []Property
}

type Array struct {
	Stmt
	Type     string
	Elements []Stmt
}

type String struct {
	Stmt
	Type  string
	Value string
}

type Int struct {
	Stmt
	Type  string
	Value int64
}

type Float struct {
	Stmt
	Type  string
	Value float64
}

type Null struct {
	Stmt
	Type  string
	Value any
}

type Boolean struct {
	Stmt
	Type  string
	Value bool
}