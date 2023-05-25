package interpreter

import (
	"strconv"
	// "Soup/src/lexer/tokens/kind"
	// "reflect"
	"Soup/src/parser"
)

type Interp interface {

	eval() RuntimeVal
	eval_binary_expr() RuntimeVal

}

type Inte struct {
	Interp
}

func (s *Inte) eval(node parser.Stmt) RuntimeVal {

	switch (node.GetType()){
		case "NumericLiteral":
			o, _ := strconv.ParseFloat(string(node.GetValue()), 64)
			return MK_NUMERAL(o)

		case "NullLiteral":
			return MK_NULL()

		case "Program":
			return MK_NULL()

		case "BinaryExpr":
			return s.eval_binary_expr(node)
	}

}