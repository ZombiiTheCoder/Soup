package runtime

import "Soup/src/parser/ast"

func Make_Null() RuntimeVal {
	return NullVal{ Val: nul }
}

func Make_Numeral(val int) RuntimeVal {
	return NumeralVal{ Val: val }
}

func Make_Float(val float64) RuntimeVal {
	return FloatVal{ Val: val }
}

func Make_Boolean(val bool) RuntimeVal {
	return BooleanVal{ Val: val }
}

func Make_String(val string) RuntimeVal {
	return StringVal{ Val: val }
}

func Make_ObjectVal(val map[string]RuntimeVal) RuntimeVal {
	return ObjectVal{ Val: val }
}

func Make_MemberVal(val map[string]RuntimeVal) RuntimeVal {
	return MemberVal{ Val: val }
}

func Make_Function(name string, call FuncCall) RuntimeVal {
	return NativeFuncVal{ Name:name, Call: call }
}

func Make_Soup_Func(name string, params []string, env Env, body []ast.Stmt) RuntimeVal {
	return FuncVal{ Name:name, Params:params, DecEnv:env, Body:body }
}

func Make_Ret_Stmt(st RuntimeVal) RuntimeVal {
	return RetVal{ Val:st }
}