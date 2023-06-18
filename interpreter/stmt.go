package interpreter

import (
	"os"
	"soup/ast"
	"soup/lexer"
	"soup/parser"
	"soup/runtime"
	"strings"
)

func (s *Interpreter) EvalProgram(node ast.Program, env runtime.Env) runtime.Val {

	var lastEval runtime.Val = runtime.Null{Value: "null", Type: "Null"}

	for _, v := range node.Body {
		lastEval = s.Eval(v, env)
	}

	return lastEval

}

func (s *Interpreter) EvalVarDec(node ast.VarDec, env runtime.Env) runtime.Val {

	if node.Value != nil {
		value := s.Eval(node.Value.(ast.Stmt), env)
		return env.DeclareVar(node.Name, value, node.NotMut)
	}
	return env.DeclareVar(node.Name, runtime.Null{Value: "null", Type: "Null"}, node.NotMut)

}

func (s *Interpreter) EvalFunction(node ast.FuncDec, env runtime.Env) runtime.Val {

	function := runtime.Func{
		Name: node.Name,
		Params: node.Params,
		DecEnv: env,
		Body: node.Body,
		Type: "Func",
	}

	env.DeclareVar(node.Name, function, true)
	return function

}

// func (s *Interpreter) EvalReturn(node ast.ReturnStmt, env runtime.Env) runtime.Val {

// 	// fmt.Println(s.Eval(node.Value, env))
// 	// v:=s.Eval(node.Value, env)
// 	// v=v.SetReturned(true)
// 	return node.Value

// }

func (s *Interpreter) EvalImport(node ast.ImpStmt, env runtime.Env) runtime.Val {

	var importReturn runtime.Val

	file, _ := os.ReadFile(node.File)

	if env.Parent != nil {
		lexers := lexer.Lexer{FileName: node.File, Chars: strings.Split(string(file), "")}
		parsed := parser.Parser{Tokens: lexers.Tokenize(), FileName: node.File, I: 0}
		intp := (Interpreter{Ast: parsed.Parse(), Env: env.Parent.(runtime.Env)})
		importReturn = intp.Eval(intp.Ast, env)
	}else {
		lexers := lexer.Lexer{FileName: node.File, Chars: strings.Split(string(file), "")}
		parsed := parser.Parser{Tokens: lexers.Tokenize(), FileName: node.File, I: 0}
		intp := (Interpreter{Ast: parsed.Parse(), Env: env})
		importReturn = intp.Eval(intp.Ast, env)
	}

	return importReturn

}

func (s *Interpreter) EvalBlock(node ast.BlockStmt, env runtime.Env) runtime.Val {

	newEnv := CreateEnvWithParent(env)

	for _, v := range node.Body {
		s.Eval(v, newEnv)
	}

	return runtime.Null{Value: "null", Type: "Null"}

}

func (s *Interpreter) EvalIf(node ast.IfStmt, env runtime.Env) runtime.Val {
	condition := s.Eval(node.Test, env)
	env2:=CreateEnvWithParent(env)
	if runtime.IsTrue(condition) {
		for _, v := range node.Consquent {
			s.Eval(v, env2)
		}
	}else if node.Alternate != nil{
		for _, v := range node.Alternate.([]ast.Stmt) {
			s.Eval(v, env2)
		}
	}

	return runtime.Null{Value: "null", Type: "Null"}

}

func (s *Interpreter) EvalWhile(node ast.WhileStmt, env runtime.Env) runtime.Val {
	condition := s.Eval(node.Test, env)
	for runtime.IsTrue(condition) {
		env2:=CreateEnvWithParent(env)
		for _, v := range node.Consquent {
			s.Eval(v, env2)
		}
		if !runtime.IsTrue(s.Eval(node.Test, env2)){
			break
		}
	}

	return runtime.Null{Value: "null", Type: "Null"}

}
