package interpreter

import (
	"soup/ast"
	"soup/runtime"
	"soup/utils"
)

type Interpreter struct {
	runtime.Interpreter
	Ast ast.Stmt
	Env runtime.Env
}

func (s *Interpreter) Eval(node ast.Stmt, env runtime.Env) runtime.Val {

	switch node.GetType() {
	case "Identifier":
		return s.EvalIdentifier(node.(ast.Identifier), env)
	case "StringLiteral":
		return runtime.String{Value: node.(ast.StringLiteral).Valu, Type: "String"}
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
		return s.EvalMemberTypes(node.(ast.MemberExpr), env)
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
	case "ReturnStmt":
		return s.EvalReturn(node.(ast.ReturnStmt), env)
	default:
		utils.Error("Unidentified AstNode Of Type %v", node)
		return runtime.Null{Value: "Null"}
	}

}