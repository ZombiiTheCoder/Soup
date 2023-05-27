package runtime

import "Soup/src/parser/ast"

type NullVal struct{
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
	Val string
}

type ObjectVal struct {
	RuntimeVal
	Val map[string]RuntimeVal
}

type MemberVal struct {
	RuntimeVal
	Val map[string]RuntimeVal
}

type FuncCall func(args []RuntimeVal, env Env) RuntimeVal

type NativeFuncVal struct {
	RuntimeVal
	Name string
	Call FuncCall
}

type FuncVal struct {
	RuntimeVal
	Name string
	Params []string
	DecEnv Env
	Body []ast.Stmt
}

type RetVal struct {
	RuntimeVal
	Val RuntimeVal
}