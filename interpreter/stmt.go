package interpreter

import (
	"os"
	"soup/ast"
	"soup/lexer"
	"soup/parser"
	"soup/runtime"
	"soup/utils"
	"strings"
)

func (s *Interpreter) EvalProgram(node ast.Program, env runtime.Env) runtime.Val {

	var lastEval runtime.Val = runtime.Null{Value: "null", Type: "Null"}

	for _, v := range node.Body {
		lastEval = s.Eval(v, env)
		
		if lastEval.GetType() == "ReturnStmt" {
			utils.Error("Toplevel Return Not Supported")
		}
	}

	return lastEval

}

func (s *Interpreter) EvalVarDec(node ast.VarDec, env runtime.Env) runtime.Val {

	var value runtime.Val = runtime.Null{Value: "null", Type: "Null"}

	if node.Value != nil {
		value = s.Eval(node.Value.(ast.Stmt), env)
	}

	return env.DeclareVar(node.Name, value, node.NotMut)

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

func (s *Interpreter) EvalReturn(node ast.ReturnStmt, env runtime.Env) runtime.Val {

	return runtime.Return{Value: s.Eval(node.Value, env), Type: "Return"}

}

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

	newEnv := runtime.CreateEnvWithParent(env)

	for _, v := range node.Body {
		s.Eval(v, newEnv)
	}

	return runtime.Null{Value: "null", Type: "Null"}

}

func (s *Interpreter) EvalIf(node ast.IfStmt, env runtime.Env) runtime.Val {
	condition := s.Eval(node.Test, env)
	if runtime.IsTrue(condition) {
		for _, v := range node.Consquent {
			s.Eval(v, env)
		}
	}else if node.Alternate != nil{
		for _, v := range node.Alternate.([]ast.Stmt) {
			s.Eval(v, env)
		}
	}

	return runtime.Null{Value: "null", Type: "Null"}

}
