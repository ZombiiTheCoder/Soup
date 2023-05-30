package runtime

import "Soup/src/ast"

type NullVal struct {
	RuntimeVal
	Val null
}

type NumeralVal struct {
	RuntimeVal
	Val int
}

type FloatVal struct {
	RuntimeVal
	Val float64
}

type BooleanVal struct {
	RuntimeVal
	Val bool
}

type StringVal struct {
	RuntimeVal
	Val         string
	ObjElements map[string]RuntimeVal
}

type ObjectVal struct {
	RuntimeVal
	ObjElements map[string]RuntimeVal
}

type MemberVal struct {
	RuntimeVal
	ObjElements map[string]RuntimeVal
}

type FuncCall func(args []RuntimeVal, env Env) RuntimeVal

type NativeFuncVal struct {
	RuntimeVal
	Name string
	Call FuncCall
}

type FuncVal struct {
	RuntimeVal
	Name   string
	Params []string
	DecEnv Env
	Body   []ast.Stmt
}

type RetVal struct {
	RuntimeVal
	Val RuntimeVal
}

type ArrayVal struct {
	RuntimeVal
	Elements    []RuntimeVal
	ObjElements map[string]RuntimeVal
}
