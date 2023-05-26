package interpreter

import (
	// "Soup/src/utils/fmt"
	f "fmt"
	"os"
	"strconv"
	// "Soup/src/lexer/tokens/kind"
	// "reflect"
	rt "Soup/src/interpreter/runtime" 
	"Soup/src/parser"
	"Soup/src/parser/ast"
)

type Interp interface {

	Eval() rt.RuntimeVal
	Eval_binary_expr() rt.RuntimeVal
	Eval_numeric_binary_expr() rt.RuntimeVal
	Eval_nonnumerical_binary_expr() rt.RuntimeVal
	Eval_program() rt.RuntimeVal

}

type Inte struct {
	Interp
}

func (s *Inte) Eval(node ast.Stmt, env Env) rt.RuntimeVal {

	switch (node.GetType()){

		case "StringLiteral":
			return rt.Make_String(node.GetValue())

		case "NumericLiteral":
			o, _ := strconv.ParseInt(string(node.GetValue()), 10, 64)
			return rt.Make_Numeral(int(o))

		case "FloatLiteral":
			o, _ := strconv.ParseFloat(string(node.GetValue()), 64)
			return rt.Make_Float(o)

		case "Identifier":
			return s.Eval_identifier(node.(ast.Identifier), env)

		case "Program":
			return s.Eval_program(node.(ast.Program), env)

		case "BinaryExpr":
			// fmt.Prints.PrintLn(node)
			return s.Eval_binary_expr(node.(ast.BinaryExpr), env)
	
		case "VarDec":
			return s.Eval_Var_dec(node.(ast.VarDec), env)

		case "ObjectLiteral":
			return s.Eval_object_expr(node.(ast.ObjectLiteral), env)

		case "MemberExpr":
            return s.Eval_member_expr(node.(ast.MemberExpr), env)

		default:
			f.Printf("This AST Node has not yet been setup for interpretation. %v\n", node)
			os.Exit(1)

	}

	return rt.Make_Null()

}

func BuildInterpreter(Src string, env Env) (rt.RuntimeVal, Env) {

	v := Inte{}
	return v.Eval(parser.BuildParser(Src), env), env

}