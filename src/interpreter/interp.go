package interpreter

import (
	// "Soup/src/utils/fmt"
	f "fmt"
	"os"
	"strconv"

	// "Soup/src/lexer/tokens/kind"
	// "reflect"
	"Soup/src/ast"
	rt "Soup/src/interpreter/runtime"
	"Soup/src/parser"
)

var StdPath string
var FilePath string
var ExeDir string

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

func (s *Inte) Eval(node ast.Stmt, env rt.Env) rt.RuntimeVal {

	switch node.GetType() {

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

	case "RelationalExpr":
		return s.Eval_Relational_Expr(node.(ast.RelationalExpr), env)

	case "UnaryExpr":
		return s.Eval_Unary_expr(node.(ast.UnaryExpr), env)

	case "BinaryExpr":
		return s.Eval_binary_expr(node.(ast.BinaryExpr), env)

	case "VarDec":
		return s.Eval_Var_dec(node.(ast.VarDec), env)

	case "FuncDec":
		return s.Eval_Func_dec(node.(ast.FuncDec), env)

	case "RetStmt":
		return s.Eval_Ret_stmt(node.(ast.RetStmt), env)

	case "ImpStmt":
		return s.Eval_Imp_stmt(node.(ast.ImpStmt), env)

	case "AssignExpr":
		return s.Eval_Assign_Expr(node.(ast.AssignExpr), env)

	case "CallExpr":
		return s.Eval_call_expr(node.(ast.CallExpr), env)

	case "ObjectLiteral":
		return s.Eval_object_expr(node.(ast.ObjectLiteral), env)

	case "MemberExpr":
		return s.Eval_member_expr(node.(ast.MemberExpr), env)

	case "BlockStmt":
		return s.Eval_Block_Stmt(node.(ast.BlockStmt), env)

	case "IfStmt":
		return s.Eval_If_Stmt(node.(ast.IfStmt), env)

	case "WhileStmt":
		return s.Eval_While_Stmt(node.(ast.WhileStmt), env)

	case "ArrayExpr":
		return s.Eval_Array_Expr(node.(ast.ArrayExpr), env)

	default:
		f.Printf("\nThis AST Node has not yet been setup for interpretation. %v Of Type %v\n", node, node.GetType())
		os.Exit(1)

	}

	return rt.Make_Null()

}

func BuildInterpreter(exeDir, stdPath, filePath, Src string, env rt.Env) (rt.RuntimeVal, rt.Env) {
	ExeDir = exeDir
	StdPath = stdPath
	FilePath = filePath
	v := Inte{}
	return v.Eval(parser.BuildParser(Src), env), env
}
