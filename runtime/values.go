package runtime

import (
	"soup/ast"
)

type Val interface {
	GetType() string
}

type Null struct {
	Val
	Type string
	Value string
}

type Int struct {
	Val
	Type string
	Value int
}


type Float struct {
	Val
	Type string
	Value float64
}


type Bool struct {
	Val
	Type string
	Value bool
}


type String struct {
	Val
	Type string
	Value          string
	ObjectElements map[string]Val
}


type Object struct {
	Val
	Type string
	ObjectElements map[string]Val
}


type Member struct {
	Val
	Type string
	ObjectElements map[string]Val
}


type NativeFuncCall func(ast_raw []ast.Expr, args []Val, scope Env) Val
type NativeMethodCall func(parent Val, ast_raw []ast.Expr, args []Val, scope Env) Val

type NativeFunc struct {
	Val
	Type string
	Name string
	Call NativeFuncCall
}

type NativeMethod struct {
	Val
	Type string
	Name string
	Call NativeMethodCall
}


type Func struct {
	Val
	Type string
	Name   string
	Params []string
	DecEnv Env
	Body   []ast.Stmt
}


type Return struct {
	Val
	Type string
	Value Val
}


type Array struct {
	Val
	Type string
	Elements []Val
	ObjectElements map[string]Val
}
