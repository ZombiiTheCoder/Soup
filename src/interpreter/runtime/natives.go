package runtime

import (
	f "fmt"
	"strings"
)

var Qenv Env

func Func(name string, funct FuncCall) func() {
	f := func() {
		Qenv.DeclareVar(name, Make_Function(name, funct), true)
	}
	return f
}

func DeclareNatives(env Env) Env {
	Qenv = env

	// DPrint()
	// DPrintLn()
	DList_Get()
	DRemove()
	DExit()
	DWait()
	DLen()
	DEFIFNOT()
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
		return FormatObject(r.(ObjectVal).ObjElements, 0)
	case "MemberVal":
		return r.(MemberVal).ObjElements
	case "RetVal":
		return GetVal(r.(RetVal).Val)
	case "NativeFuncVal":
		return f.Sprintf("fn %v (args []RuntimeVal, scope Env) RuntimeVal { Native Code }", r.(NativeFuncVal).Name)
	case "FuncVal":
		q := strings.Join(r.(FuncVal).Params, ", ")
		return f.Sprintf("fn %v (%v) { Soup Code }", r.(FuncVal).Name, q)
	case "ArrayVal":
		return FormatArray(r.(ArrayVal).Elements)
	default:
		return f.Sprintf("Value Not Avalible, Value Is Of Type %v", r.GetType())
	}

}

func IsTrue(condition RuntimeVal) bool {
	switch condition.GetType() {
	case "NullVal":
		return false
	case "NumeralVal":
		if condition.(NumeralVal).Val != 0 {
			return true
		} else {
			return false
		}
	case "BooleanVal":
		return condition.(BooleanVal).Val
	case "NativeFuncVal":
		return true
	case "FuncVal":
		return true
	case "StringVal":
		if strings.TrimSpace(condition.(StringVal).Val) != "" {
			return true
		} else {
			return false
		}

	default:
		return false
	}
}

func IsNumeric(val RuntimeVal) bool {
	switch val.GetType() {
	case "NumeralVal":
		return true
	case "FloatVal":
		return true
	default:
		return false
	}
}
