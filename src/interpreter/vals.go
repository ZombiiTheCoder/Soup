package interpreter

// import "Soup/src/parser"

type null *struct{}
var niz null

type RuntimeVal interface{}

type NullVal struct{
	RuntimeVal
	Val null
}

func MK_NULL() RuntimeVal{
	return NullVal{Val: niz}
}

type NumeralVal struct {
	RuntimeVal
	Val float64
}

func MK_NUMERAL(f float64) RuntimeVal{
	return NumeralVal{Val: f}
}

type BooleanVal struct {
	RuntimeVal
	Val bool
}

func MK_BOOL(f bool) RuntimeVal{
	return BooleanVal{Val: f}
}

type StringVal struct {
	RuntimeVal
	Val string
}

func MK_STRING(f string) RuntimeVal{
	return StringVal{Val: f}
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