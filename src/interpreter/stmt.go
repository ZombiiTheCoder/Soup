package interpreter

import (
	"Soup/src/ast"
	rt "Soup/src/interpreter/runtime"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func (s *Inte) Eval_program(prg ast.Program, env rt.Env) rt.RuntimeVal {

	lastEvaluated := rt.Make_Null()

	for _, v := range prg.Body {
		lastEvaluated = s.Eval(v, env)
		if lastEvaluated.GetType() == "RetVal" {
			fmt.Println("Return Is Not Allowed Outside Function")
		}
	}

	return lastEvaluated

}

func (s *Inte) Eval_Var_dec(v ast.VarDec, env rt.Env) rt.RuntimeVal {

	var val rt.RuntimeVal
	if v.Valu != nil {
		val = s.Eval(v.Valu.(ast.Stmt), env)
	} else {
		val = rt.Make_Null()
	}
	return env.DeclareVar(v.Idnt, val, v.Cont)

}

func (s *Inte) Eval_Func_dec(fn ast.FuncDec, env rt.Env) rt.RuntimeVal {

	funcr := rt.Make_Soup_Func(
		fn.Name,
		fn.Params,
		env,
		fn.Body,
	)

	env.DeclareVar(fn.Name, funcr, true)
	return funcr
}

func (s *Inte) Eval_Ret_stmt(ret ast.RetStmt, env rt.Env) rt.RuntimeVal {

	rtv := rt.Make_Ret_Val(
		s.Eval(ret.Valu, env),
	)

	return rtv

}

func readFile(d string) string {
	body, err := ioutil.ReadFile(d)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func (s *Inte) Eval_Imp_stmt(dec ast.ImpStmt, env rt.Env) rt.RuntimeVal {

	file := ""
	if !strings.Contains(dec.File, ".soup") {
		file = filepath.Join(dec.File, "/init.soup")
	} else {
		file = dec.File
	}

	f := ""
	if strings.Contains(dec.File, "@") {
		f = filepath.Join(StdPath, strings.ReplaceAll(file, "@", ""))
	} else {
		f = filepath.Join(FilePath, file)
	}

	var rq rt.RuntimeVal

	if env.Parent != nil {
		q, _ := BuildInterpreter(ExeDir, StdPath, filepath.Dir(f), readFile(f), env.Parent.(rt.Env))
		rq = q
	} else {
		q, _ := BuildInterpreter(ExeDir, StdPath, filepath.Dir(f), readFile(f), env)
		rq = q
	}
	return rq

}

func (s *Inte) Eval_Block_Stmt(Block ast.BlockStmt, env rt.Env) rt.RuntimeVal {

	newEnv := rt.CreateEnvWithParent(env)

	for _, v := range Block.Body {
		s.Eval(v, newEnv)
	}

	return rt.Make_Null()
}

func (s *Inte) Eval_If_Stmt(ifs ast.IfStmt, env rt.Env) rt.RuntimeVal {
	condition := s.Eval(ifs.Condition, env)
	if rt.IsTrue(condition) {
		s.Eval(ifs.Consquent, env)
	} else if ifs.Alternate != nil {
		s.Eval(ifs.Alternate.(ast.Stmt), env)
	}

	return rt.Make_Null()
}

func (s *Inte) Eval_While_Stmt(ifs ast.WhileStmt, env rt.Env) rt.RuntimeVal {
	condition := s.Eval(ifs.Condition, env)

	for rt.IsTrue(condition) {
		s.Eval(ifs.Consquent, env)
		condition = s.Eval(ifs.Condition, env)
	}

	return rt.Make_Null()
}
