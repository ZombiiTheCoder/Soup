package interpreter

import (
	"fmt"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
)

type Interpreter struct {
	runtime.RTInterpreter
	Ast ast.Stmt
	Env runtime.Env
}

func (s *Interpreter) Eval(node ast.Stmt, env runtime.Env) runtime.Val {

	switch node.GetType() {
	case "Identifier":
		return s.EvalIdentifier(node.(ast.Identifier), env)
	case "StringLiteral":

		string_vals := make(map[string]runtime.Val)
		string_vals["length"] = runtime.Int{Type: "Int", Value: len(node.(ast.StringLiteral).Valu)}
		string_vals["format"] = runtime.NativeFunc{Type: "NativeFunc", Name:"format", Call: func(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {
			nargs := make([]any, 0)
			for _, v := range args {
				switch v.GetType() {
				case "Int":
					nargs = append(nargs, fmt.Sprintf("%v", v.(runtime.Int).Value))
				case "String":
					nargs = append(nargs, v.(runtime.String).Value)
				case "Bool":
					nargs = append(nargs, fmt.Sprintf("%v", v.(runtime.Bool).Value))
				case "Null":
					nargs = append(nargs, v.(runtime.Null).Value)
				default:
					nargs = append(nargs, fmt.Sprintf("%v", v))
				}
			}
			return s.Eval(ast.StringLiteral{
				Type: "StringLiteral",
				Valu: fmt.Sprintf(node.(ast.StringLiteral).Valu, nargs...),
			}, scope)
		}}

		return runtime.String{
			Value: node.(ast.StringLiteral).Valu, 
			Type: "String",
			ObjectElements: string_vals,
		}
	case "IntegerLiteral":
		return runtime.Int{Value: int(node.(ast.IntegerLiteral).Valu), Type: "Int"}
	case "FloatLiteral":
		return runtime.Float{Value: node.(ast.FloatLiteral).Valu, Type: "Float"}
	case "NullLiteral":
		return runtime.Null{Value: node.(ast.NullLiteral).Valu, Type: "Null"}
	case "ObjectLiteral":
		return s.EvalObject(node.(ast.ObjectLiteral), env)
	case "AssignExpr":
		return s.EvalAssign(node.(ast.AssignExpr), env)
	case "TernaryExpr":
		return s.EvalTernary(node.(ast.TernaryExpr), env)
	case "LogicalExpr":
		return s.EvalLogical(node.(ast.LogicalExpr), env)
	case "BinaryExpr":
		return s.EvalBinaryTypes(node.(ast.BinaryExpr), env)
	case "UnaryExpr":
		return s.EvalUnary(node.(ast.UnaryExpr), env)
	case "CallExpr":
		return s.EvalCall(node.(ast.CallExpr), env)
	case "MemberExpr":
		return s.EvalMember(node.(ast.MemberExpr), env)
	case "ArrayExpr":
		return s.EvalArray(node.(ast.ArrayExpr), env)
	case "Property":
		return s.EvalProperty(node.(ast.Property), env)
	case "VarDec":
		return s.EvalVarDec(node.(ast.VarDec), env)
	case "Program":
		return s.EvalProgram(node.(ast.Program), env)
	case "FuncDec":
		return s.EvalFunction(node.(ast.FuncDec), env)
	case "ImpStmt":
		return s.EvalImport(node.(ast.ImpStmt), env)
	case "IfStmt":
		return s.EvalIf(node.(ast.IfStmt), env)
	case "WhileStmt":
		return s.EvalWhile(node.(ast.WhileStmt), env)
	// case "ReturnStmt":
		// return s.EvalReturn(node.(ast.ReturnStmt), env)
	default:
		utils.Error("Unidentified AstNode Of Type %v", node)
		return runtime.Null{Value: "Null"}
	}

}