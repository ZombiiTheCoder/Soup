package runtime

import (
	"soup/utils"
	"strings"
)

type Env struct {
	Parent any
	Vars   map[string]Val
	Consts map[string]bool
}

//   Syntax Errors      Type Check Errors
// 		|					|
// Lexer -> Parser -> AST -> Sematic Analyzer -> IR -> Executed

func (s Env) DeclareVar(VarName string, Value Val, cons bool) Val {

	if _, ok := s.Vars[VarName]; ok {
		utils.Error("Cannot Declare Variable %v. Variable is already declared", VarName)
	}

	s.Vars[VarName] = Value
	if cons {
		s.Consts[VarName] = true
	}

	return Value

}

func (s Env) AssignVar(VarName string, Value Val) Val {
	env := s.Resolve(VarName)
	if s.Consts[VarName] {
		utils.Error("%v Cannot Reassign Constant", VarName)
	}
	env.Vars[VarName] = Value
	return Value
}

func (s Env) LookUpVar(VarName string) Val {
	env := s.Resolve(VarName)
	return env.Vars[VarName]
}

func (s Env) Resolve(VarName string) Env {

	if _, ok := s.Vars[VarName]; ok {
		return s
	}

	if s.Parent == nil {
		utils.Error("%v Does Not Exist", VarName)
	}

	return s.Parent.(Env).Resolve(VarName)

}

func CreateEnv() Env {
	env := Env{}
	env.Vars = make(map[string]Val)
	env.Consts = make(map[string]bool)
	env.DeclareVar("true", Bool{Value: true, Type: "Bool"}, true)
	env.DeclareVar("false", Bool{Value: false, Type: "Bool"}, true)
	env = Declare(env)
	return env
}

func CreateEnvWithParent(parent Env) Env {
	env := CreateEnv()
	env.Parent = parent
	return env
}

func IsTrue(val Val) bool{
	switch val.GetType() {
	case "Null":
		return false
	case "Int":
		if (val.(Int).Value == 0){
			return false
		}else {
			return true
		}
	case "Bool":
		return val.(Bool).Value
	case "String":
		if strings.TrimSpace(val.(String).Value) != ""{
			return true
		}else {
			return false
		}
	case "Object":
		return true
	case "Member":
		return true
	case "NativeFunc":
		return true
	case "Func":
		return true
	case "Return":
		return IsTrue(val.(Return).Value)
	case "Array":
		return true
	default:
		return false
	}
}