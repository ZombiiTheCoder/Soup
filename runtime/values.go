package runtime

import "soup/ast"

type Val interface {
	GetType() string
}

type Interpreter interface {
	Eval(node ast.Node, env Env) Val

	EvalProgram(node ast.Program, env Env) Val

	EvalUnary(node ast.UnaryExpr, env Env) Val
	EvalTernary(node ast.TernaryExpr, env Env) Val
	EvalLogical(node ast.LogicalExpr, env Env) Val

	EvalBinaryTypes(node ast.BinaryExpr, env Env) Val
	EvalBinaryBitwise(node ast.BinaryExpr, env Env) Val
	EvalBinaryEquality(node ast.BinaryExpr, env Env) Val
	EvalBinaryRelational(node ast.BinaryExpr, env Env) Val
	
	EvalAssign(node ast.AssignExpr, env Env) Val
	EvalVarDec(node ast.VarDec, env Env) Val
	EvalFunction(node ast.FuncDec, env Env) Val
	EvalReturn(node ast.ReturnStmt, env Env) Val
	EvalImport(node ast.ImpStmt, env Env) Val

	EvalCall(node ast.Node, env Env) Val
	EvalObject(node ast.Node, env Env) Val

	EvalMemberTypes(node ast.MemberExpr, env Env) Val
	EvalMember(node ast.MemberExpr, env Env) Val
	EvalMemberComputed(node ast.MemberExpr, env Env) Val
	EvalArrayObj(node ast.MemberExpr, env Env) Val
	EvalStringObj(node ast.MemberExpr, env Env) Val

	EvalBlock(node ast.BlockStmt, env Env) Val
	EvalIf(node ast.IfStmt, env Env) Val
	EvalWhile(node ast.WhileStmt, env Env) Val
	EvalArray(node ast.ArrayExpr, env Env) Val
	EvalProperty(node ast.Property, env Env) Val

	EvalIdentifier(node ast.Identifier, env Env) Val
}

type Null struct {
	Val
	Type string
	Value string
}

type Int struct {
	Val
	Type string
	Value int
}

type Float struct {
	Val
	Type string
	Value float64
}

type Bool struct {
	Val
	Type string
	Value bool
}

type String struct {
	Val
	Type string
	Value          string
	ObjectElements map[string]Val
}

type Object struct {
	Val
	Type string
	ObjectElements map[string]Val
}

type Member struct {
	Val
	Type string
	ObjectElements map[string]Val
}

type NativeFuncCall func(args []Val, scope Env) Val

type NativeFunc struct {
	Val
	Type string
	Name string
	Call NativeFuncCall
}

type Func struct {
	Val
	Type string
	Name   string
	Params []string
	DecEnv Env
	Body   []ast.Stmt
}

type Return struct {
	Val
	Type string
	Value Val
}

type Array struct {
	Val
	Type string
	Elements []Val
	ObjectElements map[string]Val
}