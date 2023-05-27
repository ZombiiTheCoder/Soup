package runtime

import f"fmt"
import "strings"

var Qenv Env

func Func(name string, funct FuncCall) (func()) {
	f:=func(){
		Qenv.DeclareVar(name, Make_Function(name, funct), true)
	}
	return f
}

func DeclareNatives(env Env) Env {
	Qenv = env
	
	DPrint()
	DGoEval()

	return env

}

func GetVal(r RuntimeVal) any {

	switch r.GetType() {
		case "NullVal":
			return nil
		case "NumeralVal":
			return r.(NumeralVal).Val
		case "FloatVal":
			return r.(FloatVal).Val
		case "BooleanVal":
			return r.(BooleanVal).Val
		case "StringVal":
			return r.(StringVal).Val
		case "ObjectVal":
			return r.(ObjectVal).Val
		case "MemberVal":
			return r.(MemberVal).Val
		case "RetVal":
			return GetVal(r.(RetVal).Val)
		case "NativeFuncVal":
			return f.Sprintf("fn %v (args []RuntimeVal, scope Env) RuntimeVal {\n Native Code \n}", r.(NativeFuncVal).Name)
		case "FuncVal":
			q:=strings.Join(r.(FuncVal).Params, ", ")
			return f.Sprintf("fn %v (%v) {\n Soup Code \n}", r.(FuncVal).Name, q)

		default:
			return f.Sprintf("Value Not Avalible, Value Is Of Type %v", r.GetType())
	}

}