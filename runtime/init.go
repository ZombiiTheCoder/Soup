package runtime

import (
	"soup/ast"
	"strings"
)

type RTInterpreter interface {
	Eval(node ast.Node, env Env) Val

	EvalProgram(node ast.Program, env Env) Val

	EvalUnary(node ast.UnaryExpr, env Env) Val
	EvalTernary(node ast.TernaryExpr, env Env) Val
	EvalLogical(node ast.LogicalExpr, env Env) Val

	EvalBinaryTypes(node ast.BinaryExpr, env Env) Val
	EvalBinaryBitwise(left ast.Expr, right ast.Expr, op string, env Env) Val
	EvalBinaryAdditive(left ast.Expr, right ast.Expr, op string, env Env) Val
	EvalBinaryEquality(left ast.Expr, right ast.Expr, op string, env Env) Val
	EvalBinaryRelational(left ast.Expr, right ast.Expr, op string, env Env) Val
	EvalBinaryConcatenation(left ast.Expr, right ast.Expr, op string, env Env) Val
	
	EvalAssign(node ast.AssignExpr, env Env) Val
	EvalVarDec(node ast.VarDec, env Env) Val
	EvalFunction(node ast.FuncDec, env Env) Val
	EvalReturn(node ast.ReturnStmt, env Env) Val
	EvalImport(node ast.ImpStmt, env Env) Val

	EvalCall(node ast.Node, env Env) Val
	EvalObject(node ast.Node, env Env) Val

	EvalMemberTypes(node ast.MemberExpr, env Env) Val
	EvalMember(node ast.MemberExpr, env Env) Val
	EvalArrayObj(node ast.MemberExpr, env Env) Val
	FixStringObj(node String, keys []string, env Env) Val

	EvalBlock(node ast.BlockStmt, env Env) Val
	EvalIf(node ast.IfStmt, env Env) Val
	EvalWhile(node ast.WhileStmt, env Env) Val
	EvalArray(node ast.ArrayExpr, env Env) Val
	EvalProperty(node ast.Property, env Env) Val

	EvalIdentifier(node ast.Identifier, env Env) Val
}

func IsTrue(val Val) bool{
	switch val.GetType() {
	case "Null":
		return false
	case "Int":
		if (val.(Int).Value == 0){
			return false
		}else {
			return true
		}
	case "Bool":
		return val.(Bool).Value
	case "String":
		if strings.TrimSpace(val.(String).Value) != ""{
			return true
		}else {
			return false
		}
	case "Object":
		return true
	case "Member":
		return true
	case "NativeFunc":
		return true
	case "Func":
		return true
	case "Return":
		return IsTrue(val.(Return).Value)
	case "Array":
		return true
	default:
		return false
	}
}