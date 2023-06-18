package interpreter

import (
	"os"
	"soup/ast"
	"soup/lexer"
	"soup/parser"
	"soup/runtime"
	"soup/runtime/native_functions"
	"soup/utils"
	"strings"
)

func Func_Eval(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	if len(args) > 1 || len(args) < 1 {
		utils.Error("There must be one or more arguments in type(any) function")
	}

	if args[0].GetType() != "String" {
		utils.Error("Contents must be string for eval(string) function")
	}

	Contents := args[0].(runtime.String).Value
	Chars := strings.Split(Contents, "")
	Lexer := lexer.Lexer{FileName: "eval.soup", Chars: Chars, I: 0, Line: 1, Column: 1}

	Parser := parser.Parser{ I: 0, Tokens: Lexer.Tokenize(), FileName: "eval.soup"}

	type VInterpreter struct {
		Interpreter
		Ast ast.Stmt
		Env runtime.Env
	}

	Interp := VInterpreter{}
	Interp.Ast = Parser.Parse()
	Interp.Env = CreateEnv()
	return Interp.Eval(Interp.Ast, Interp.Env)

}

func Func_Type(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	if len(args) > 1 || len(args) < 1 {
		utils.Error("There must be one argument in type(any) function")
	}

	type VInterpreter struct {
		Interpreter
		Ast ast.Stmt
		Env runtime.Env
	}

	Interp := VInterpreter{}
	Interp.Env = CreateEnv()
	return Interp.Eval(ast.StringLiteral{Type: "StringLiteral", Valu: args[0].GetType()}, Interp.Env)

}

func Func_Exit(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	if len(args) > 1 || len(args) < 1 {
		utils.Error("There must be one argument in type(int) function")
	}

	os.Exit(args[0].(runtime.Int).Value)
	
	return runtime.Null{Type: "Null", Value: "null"}

}

func Func_Append(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	if args[0].GetType() != "Array" {
		utils.Error("The First Argument must be Array for eval(array) function")
	}

	elems := args[1:]

	array := args[0].(runtime.Array)
	array.Elements = append(array.Elements, elems...) 
	array.ObjectElements["length"] = runtime.Int{Type: "Int", Value: len(array.Elements)}

	
	return array

}

func CreateEnv() runtime.Env {

	env := runtime.Env{}
	env.Vars = make(map[string]runtime.Val)
	env.Consts = make(map[string]bool)
	env.DeclareVar("true", runtime.Bool{Value: true, Type: "Bool"}, true)
	env.DeclareVar("false", runtime.Bool{Value: false, Type: "Bool"}, true)
	env.DeclareVar("null", runtime.Null{Value: "null", Type: "Null"}, true)
	env.DeclareVar("eval", runtime.NativeFunc{Type: "NativeFunc", Name: "eval", Call: Func_Eval}, true)
	env.DeclareVar("type", runtime.NativeFunc{Type: "NativeFunc", Name: "type", Call: Func_Type}, true)
	env.DeclareVar("exit", runtime.NativeFunc{Type: "NativeFunc", Name: "exit", Call: Func_Exit}, true)
	env.DeclareVar("append", runtime.NativeFunc{Type: "NativeFunc", Name: "append", Call: Func_Append}, true)
	env = native_functions.DeclareNatives(env)
	return env
}

func CreateEnvWithParent(parent runtime.Env) runtime.Env {
	env := CreateEnv()
	env.Parent = parent
	return env
}