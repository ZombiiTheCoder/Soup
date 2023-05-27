package runtime

import (
	f "fmt"
	"os"
)

type Envi interface{}

type Env struct{

	Envi
	Parent any
	Vars map[string]RuntimeVal
	Consts map[string]bool

}
    //   Syntax Errors      Type Check Errors
	// 		|					|
	// Lexer -> Parser -> AST -> Sematic Analyzer -> IR -> Executed

func (s Env) DeclareVar (VarName string, Val RuntimeVal, cons bool) RuntimeVal{

	if _, ok := s.Vars[VarName]; ok {
		f.Printf("\nCannot Declare Variable %v. Variable is already declared\n", VarName)
		os.Exit(1)
	}

	s.Vars[VarName] = Val
	if (cons){
		s.Consts[VarName] = true
	}

	return Val;

}

func (s Env) AssignVar (VarName string, Val RuntimeVal) RuntimeVal {
	env := s.Resolve(VarName)
	if (s.Consts[VarName]){
		f.Printf("\n%v Cannot Reassign Constant\n", VarName)
		os.Exit(1)
	}
	env.Vars[VarName] = Val
	return Val
}

func (s Env) LookUpVar (VarName string) RuntimeVal {
	env := s.Resolve(VarName)
	return env.Vars[VarName]
}

func (s Env) Resolve (VarName string) Env {
	
	if _, ok := s.Vars[VarName]; ok {
		return s
	}

	if s.Parent == nil {
		f.Printf("\n%v Does Not Exist\n", VarName)
		os.Exit(1)
	}

	return s.Parent.(Env).Resolve(VarName)

}

func CreateEnv() Env {
	env:=Env{}
	env.Vars = make(map[string]RuntimeVal )
	env.Consts = make(map[string]bool)
	env.DeclareVar("true", Make_Boolean(true), true)
	env.DeclareVar("false", Make_Boolean(false), true)
	env=DeclareNatives(env)
	return env
}

func CreateEnvWithParent(parent Env) Env {
	env:=CreateEnv()
	env.Parent = parent
	return env
}