package interpreter

import (
	"strings"
	"path/filepath"
	"fmt"
	"Soup/src/parser/ast"
	rt "Soup/src/interpreter/runtime" 
	"io/ioutil"
	"log"
	"path"
)

func (s *Inte) Eval_program (prg ast.Program, env rt.Env) rt.RuntimeVal {

	lastEvaluated := rt.Make_Null()

	for _, v := range prg.Body {
		lastEvaluated = s.Eval(v, env)
	}

	return lastEvaluated

}

func (s *Inte) Eval_Var_dec (v ast.VarDec, env rt.Env) rt.RuntimeVal {

	var val rt.RuntimeVal
	if (v.Valu != nil){
		val = s.Eval(v.Valu.(ast.Stmt), env)
	}else {
		val = rt.Make_Null()
	}

	return env.DeclareVar(v.Idnt, val, v.Cont)

}

func (s *Inte) Eval_Func_dec (fn ast.FuncDec, env rt.Env) rt.RuntimeVal {
	
	funcr := rt.Make_Soup_Func(
		fn.Name,
		fn.Params,
		env,
		fn.Body,
	)

	fmt.Println(fn.Name)
	env.DeclareVar(fn.Name, funcr, true)
	return funcr
}

func (s *Inte) Eval_Ret_stmt (ret ast.RetStmt, env rt.Env) rt.RuntimeVal {

	rtv := rt.Make_Ret_Stmt(
		s.Eval(ret.Valu, env),
	)

	return rtv

}

func readFile(d string) (string) {
	body, err := ioutil.ReadFile(d)
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
	return string(body)
}

func (s *Inte) Eval_Imp_stmt (dec ast.ImpStmt, env rt.Env) rt.RuntimeVal {

	file:=""
	if (!strings.Contains(dec.File, ".soup")){
		file=path.Join(dec.File, "/init.soup")
	}else{
		file=dec.File
	}
	
	f:=""
	if (strings.Contains(dec.File, "@")){
		f=path.Join(StdPath, strings.ReplaceAll(file, "@", ""))
	}else{
		f=path.Join(FilePath, file)
	}
	q, e := BuildInterpreter(ExeDir, StdPath, filepath.Dir(f), readFile(f), env)
	env = e
	return q

}