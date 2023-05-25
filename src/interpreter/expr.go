package interpreter

import "Soup/src/parser"

func (s *Inte) eval_binary_expr(node parser.BinaryExpr) RuntimeVal {

	left := s.eval(node.Left)
	right := s.eval(node.Right)

}