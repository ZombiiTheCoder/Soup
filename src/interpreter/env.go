package interpreter

import (
	rt "Soup/src/interpreter/runtime"
	f "fmt"
	"os"
)

type Envi interface{

	DeclareVar(VarName string, Val rt.RuntimeVal) rt.RuntimeVal
	AssignVar(VarName string, Val rt.RuntimeVal) rt.RuntimeVal
	LookupVar(VarName string) rt.RuntimeVal
	Resolve(VarName string) rt.RuntimeVal
	
}

type Env struct{

	Envi
	Parent any
	Vars map[string]rt.RuntimeVal
	Consts map[string]bool

}
    //   Syntax Errors      Type Check Errors
	// 		|					|
	// Lexer -> Parser -> AST -> Sematic Analyzer -> IR -> Executed

func (s Env) DeclareVar (VarName string, Val rt.RuntimeVal, cons bool) rt.RuntimeVal{

	if _, ok := s.Vars[VarName]; ok {
		f.Printf("Cannot Declare Variable %v. Variable is already declared\n", VarName)
		os.Exit(1)
	}

	if (cons){
		s.Consts[VarName] = true
	}

	s.Vars[VarName] = Val

	return Val;

}

func (s Env) AssignVar (VarName string, Val rt.RuntimeVal) rt.RuntimeVal {
	env := s.Resolve(VarName)
	env.Vars[VarName] = Val
	return Val
}

func (s Env) LookUpVar (VarName string) rt.RuntimeVal {
	env := s.Resolve(VarName)
	return env.Vars[VarName]
}

func (s Env) Resolve (VarName string) Env {

	if _, ok := s.Vars[VarName]; ok {
		return s
	}

	if s.Parent == nil {
		f.Printf("%v Does Not Exist\n", VarName)
		os.Exit(1)
	}

	if (s.Consts[VarName]){
		f.Printf("%v Cannot Reassign Constant\n", VarName)
		os.Exit(1)
	}

	return s.Parent.(Env).Resolve(VarName)

}

func CreateEnv() Env {
	env:=Env{}
	env.Vars = make(map[string]rt.RuntimeVal)
	env.Consts = make(map[string]bool)
	env.DeclareVar("true", rt.Make_Boolean(true), true)
	env.DeclareVar("false", rt.Make_Boolean(false), true)
	return env
}