package interpreter

import (
	"Soup/src/utils/fmt"
	"strconv"
	// "Soup/src/lexer/tokens/kind"
	// "reflect"
	"Soup/src/parser"
)

type Interp interface {

	Eval() RuntimeVal
	Eval_binary_expr() RuntimeVal
	Eval_numeric_binary_expr() RuntimeVal
	Eval_program() RuntimeVal

}

type Inte struct {
	Interp
}

func (s *Inte) Eval(node parser.Stmt) RuntimeVal {

	switch (node.GetType()){
		case "NumericLiteral":
			o, _ := strconv.ParseFloat(string(node.GetValue()), 64)
			return MK_NUMERAL(o)

		case "NullLiteral":
			return MK_NULL()

		case "Program":
			return s.Eval_program(node.(parser.Program))

		case "BinaryExpr":
			// fmt.Prints.PrintLn(node)
			return s.Eval_binary_expr(node.(parser.BinaryExpr))
	
		default:
			fmt.Prints.ErrorF(
				"This AST Node has not yet been setup for interpretation. %v",
				node,
			)

	}

	return MK_NULL()

}