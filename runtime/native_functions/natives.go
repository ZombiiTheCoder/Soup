package native_functions

import (
	"soup/runtime"
)

func DeclareNatives(env runtime.Env) runtime.Env {
	env.DeclareVar("print", runtime.NativeFunc{Type: "NativeFunc", Name: "print", Call: Func_Print}, true)
	env.DeclareVar("raw_print", runtime.NativeFunc{Type: "NativeFunc", Name: "raw_print", Call: Func_RawPrint}, true)

	return env
}


func getVal(v runtime.Val) any {
	if v == nil{
		return v
	}
	switch v.GetType() {
	case "Int":
		return v.(runtime.Int).Value
	case "Float":
		return v.(runtime.Float).Value
	case "String":
		return v.(runtime.String).Value
	case "Bool":
		return v.(runtime.Bool).Value
	case "Null":
		return v.(runtime.Null).Value
	}
	return v

}