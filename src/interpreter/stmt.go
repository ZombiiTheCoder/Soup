package interpreter

import "Soup/src/parser/ast"

func (s *Inte) Eval_program (prg ast.Program) RuntimeVal {

	lastEvaluated := MK_NULL()

	for _, v := range prg.Body {
		lastEvaluated = s.Eval(v)
	}

	return lastEvaluated

}