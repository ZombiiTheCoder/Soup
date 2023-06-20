package builder

import (
	"soup/ast"
	"soup/interpreter"
	"soup/lexer"
	"soup/parser"
	"soup/runtime"
	"soup/tokens"
	"soup/utils"
	"strings"
)

func CreateNewLexer(File string) []tokens.Token {

	Contents := utils.ReadFile(File)
	Chars := strings.Split(Contents, "")
	Lexer := lexer.Lexer{FileName: File, Chars: Chars, I: 0, Line: 1, Column: 1}

	return Lexer.Tokenize()

}

func CreateNewParser(File string) ast.Stmt {

	Parser := parser.Parser{ I: 0, Tokens: CreateNewLexer(File), FileName: File}
	return Parser.Parse()

}

func CreateNewIntepreter(File string, env interface{}) runtime.Val {
	cenv := interpreter.CreateEnv()
	if env != nil {
		cenv = env.(runtime.Env)
	}
	Interpreter := interpreter.Interpreter{}
	Interpreter.Ast = CreateNewParser(File)
	Interpreter.Env = cenv
	return Interpreter.Eval(Interpreter.Ast, Interpreter.Env)
}

func CreateNewIntepreterWithText(Text string, env interface{}) runtime.Val {
	cenv := interpreter.CreateEnv()
	if env != nil {
		cenv = env.(runtime.Env)
	}
	Chars := strings.Split(Text, "")
	Lexer := lexer.Lexer{FileName: "console.soup", Chars: Chars, I: 0, Line: 1, Column: 1}
	Parser := parser.Parser{ I: 0, Tokens: Lexer.Tokenize(), FileName: "console.soup"}
	Interpreter := interpreter.Interpreter{}
	Interpreter.Ast = Parser.Parse()
	Interpreter.Env = cenv
	return Interpreter.Eval(Interpreter.Ast, Interpreter.Env)
}