package parser

import (
	"fmt"
	"os"
	"soup/lexer"
	. "soup/parser/ast"
)

type Parser struct {
	pointer int
	tokens  []lexer.Token
}

func (s *Parser) notEof() bool {
	return s.tokens[s.pointer].Type == lexer.EOF
}
// func (s *Parser) peek(p int) lexer.Token {
// 	if (s.pointer + p) >= len(s.tokens) {
// 		return lexer.Token{Value: "Unexpected End Of File", Type: lexer.EOF}
// 	}
// 	return s.tokens[s.pointer+p]
// }

func (s *Parser) at() lexer.Token {
	if (s.pointer) >= len(s.tokens) {
		return lexer.Token{Value: "Unexpected End Of File", Type: lexer.EOF}
	}
	return s.tokens[s.pointer]
}

func (s *Parser) next() lexer.Token {
	if s.pointer+1 >= len(s.tokens) {
		return s.tokens[len(s.tokens)-1]
	}
	s.pointer++
	return s.tokens[s.pointer-1]
}

func (s *Parser) expect(Type lexer.Types) lexer.Token {
	if s.pointer+1 >= len(s.tokens) {
		return s.tokens[len(s.tokens)-1]
	}
	s.pointer++
	last := s.tokens[s.pointer-1]
	if last.Type != Type {
		fmt.Println("Unexpected Token", last, "Expecting Token Type",Type, "Line:", last.Line, "Pos:", last.Column)
		os.Exit(1)
	}
	return last
}

func (s *Parser) eatSemi() {
	if s.at().Type == lexer.SemiColon {
		s.next()
	}
}

func (s *Parser) Init(tokens []lexer.Token) {
	s.tokens = append(s.tokens, tokens...)
	s.pointer = 0
}

func (s *Parser) Parse() Stmt {
	prg := Program{
		NodeType: "Program",
		Body: make([]Stmt, 0),
	}

	for s.notEof() {
		prg.Body = append(prg.Body, s.parseStmt())
	}

	return prg
}