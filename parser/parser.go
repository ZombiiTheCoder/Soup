package parser

import (
	"soup/ast"
	"soup/tokens"
	"soup/utils"
)

type Parser struct {
	I      int
	Tokens []tokens.Token
	FileName string
}

func (s *Parser) Peek() tokens.Token {

	if (s.I+1 >= len(s.Tokens)) {
		return s.Tokens[len(s.Tokens)-1]
	}

	return s.Tokens[s.I+1]

}

func (s *Parser) Current() tokens.Token {

	return s.Tokens[s.I]

}

func (s *Parser) Eat() tokens.Token  {
	
	if (s.I+1 >= len(s.Tokens)) {
		return s.Tokens[len(s.Tokens)-1]
	}
	
	s.I++
	return s.Tokens[s.I-1]

}

func (s *Parser) Expect(ept tokens.TokenType, err string) tokens.Token  {
	
	if (s.I+1 >= len(s.Tokens)) {
		return s.Tokens[len(s.Tokens)-1]
	}
	
	s.I++

	if s.Tokens[s.I-1].Type != ept {
		utils.Error(
			"%v\nExpected Token %v Found Token %v at line %v column %v\n Token Location %v:%v:%v",
			err,
			ept,
			s.Tokens[s.I-1].Type,
			s.Tokens[s.I-1].Line,
			s.Tokens[s.I-1].Column,
			s.Tokens[s.I-1].FileName,
			s.Tokens[s.I-1].Line,
			s.Tokens[s.I-1].Column,
		)
	}

	return s.Tokens[s.I-1]

}

func (s *Parser) EatSemi () {
	if s.Current().Type == tokens.Semicolon {
		s.Eat()
	}
}

func (s *Parser) NotEof () bool {
	return s.Tokens[s.I].Type != tokens.EndOfFile
}

func (s *Parser) Parse () ast.Stmt {
	
	program := ast.Program{Body: make([]ast.Stmt, 0), Type: "Program"}

	for s.NotEof() {
		program.Body = append(program.Body, s.ParseStmt())
	}

	return program
}