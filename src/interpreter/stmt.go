package interpreter

import "Soup/src/parser"

func (s *Inte) Eval_program (prg parser.Program) RuntimeVal {

	lastEvaluated := MK_NULL()

	for _, v := range prg.Body {
		lastEvaluated = s.Eval(v)
	}

	return lastEvaluated

}