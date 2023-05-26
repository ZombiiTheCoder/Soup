package interpreter

import "Soup/src/parser/ast"
import rt "Soup/src/interpreter/runtime" 

func (s *Inte) Eval_program (prg ast.Program, env Env) rt.RuntimeVal {

	lastEvaluated := rt.Make_Null()

	for _, v := range prg.Body {
		lastEvaluated = s.Eval(v, env)
	}

	return lastEvaluated

}

func (s *Inte) Eval_Var_dec (v ast.VarDec, env Env) rt.RuntimeVal {

	var val rt.RuntimeVal
	if (v.Valu != nil){
		val = s.Eval(v.Valu.(ast.Stmt), env)
	}else {
		val = rt.Make_Null()
	}

	return env.DeclareVar(v.Idnt, val, v.Cont)

}