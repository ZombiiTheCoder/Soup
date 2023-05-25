package interpreter

import (
	// "Soup/src/utils/fmt"
	"Soup/src/parser"
	"math"
)

func (s *Inte) Eval_binary_expr(node parser.BinaryExpr) RuntimeVal {

	left := s.Eval(node.Left)
	right := s.Eval(node.Right)

	if (left.GetType() == "NumeralVal" && right.GetType() == "NumeralVal") {
		return s.Eval_numeric_binary_expr(
			left.(NumeralVal),
			right.(NumeralVal),
			node.Op,
		)
	}

	return MK_NULL()

}

func (s *Inte) Eval_numeric_binary_expr(left, right NumeralVal, op string) RuntimeVal {

	var result float64

	switch op {
		case "+":
			result=left.Val + right.Val
		case "-":
			result=left.Val - right.Val
		case "/":
			result=left.Val / right.Val
		case "*":
			result=left.Val * right.Val
		case "%":
			result=math.Mod(left.Val, right.Val)

	}

	return MK_NUMERAL(result)

}