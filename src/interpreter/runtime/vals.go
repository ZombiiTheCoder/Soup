package runtime

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

type FuncCall func()

type NativeFuncVal struct {
	RuntimeVal
	Call FuncCall
}

// type FuncValue struct {
// 	RuntimeVal
// 	name string
// 	params []string
// 	decEnv Env
// 	body []parser.Stmt
// }