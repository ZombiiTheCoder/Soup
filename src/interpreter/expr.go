package interpreter

import (
	// "Soup/src/utils/fmt"
	f "fmt"
	"Soup/src/parser/ast"
	"math"
)

func (s *Inte) Eval_binary_expr(node ast.BinaryExpr) RuntimeVal {

	left := s.Eval(node.Left)
	right := s.Eval(node.Right)

	if (left.GetType() == "NumeralVal" && right.GetType() == "NumeralVal") {
		return s.Eval_numeric_binary_expr(
			left.(NumeralVal),
			right.(NumeralVal),
			node.Op,
		)
	}else {
		return s.Eval_nonnumerical_binary_expr(
			left,
			right,
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

func (s *Inte) Eval_nonnumerical_binary_expr(left, right RuntimeVal, op string) RuntimeVal {

	var result string
	var l any
	var r any

	switch left.GetType(){
		case "NumeralVal":
			l=left.(NumeralVal).Val
		
		case "StringVal":
			l=left.(StringVal).Val

		case "BooleanVal":
			l=left.(BooleanVal).Val

		case "NullVal":
			l="null"
	}

	switch right.GetType(){
	case "NumeralVal":
		r=right.(NumeralVal).Val
	
	case "StringVal":
		r=right.(StringVal).Val

	case "BooleanVal":
		r=right.(BooleanVal).Val

	case "NullVal":
		r="null"
}

	switch op {
		case "+":
			result=f.Sprint(l) + f.Sprint(r)

	}

	return MK_STRING(result)

}