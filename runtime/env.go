package runtime

import (
	"soup/utils"
)

type Env struct {
	Parent any
	Vars   map[string]Val
	Consts map[string]bool
}

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
		utils.Error(`Varible or Function/Method "%v" Does Not Exist`, VarName)
	}

	return s.Parent.(Env).Resolve(VarName)

}