package builder

import (
	"soup/ast"
	"soup/lexer"
	"soup/parser"
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