package lex2

import (
	"Soup/src/lexer"
	"Soup/src/utils/fmt"
	"Soup/src/lexer/tokens"
	"Soup/src/lexer/tokens/kind"
)

type Lex interface{
	At() string
	Next() int
	BuildToken() int
	Tokenize() []tokens.Token
}

type Lexer struct{
	Lex
	Chars []string
	Ip int
	Tokens []tokens.Token
}

func (s *Lexer) At() string {
	return s.Chars[s.Ip]
}

func (s *Lexer) Next() int {
	s.Ip = s.Ip + 1
	return 0
}

func (s *Lexer) BuildToken(value string, typee kind.TokenKind) int {

	s.Tokens = append(s.Tokens, tokens.Token{Value: value, Type: typee})
	return 0

}

func (s *Lexer) Tokenize() []tokens.Token {

	for (s.Ip < len(s.Chars)){
		
		println(s.Ip)
		
		if (lexer.OCT(s.At())){
			s.BuildToken(s.At(), lexer.TKNS(s.At()))
			s.Next()
		}else{
		
			if (lexer.IsNum(s.At())){
				num:=""
				for (lexer.IsNum(s.At())){
					num+=s.At()
					s.Next()
				}
				
				s.BuildToken(num, kind.Numeral)
			}
			
			if (lexer.IsSym(s.At())){
				sym:=""
				for (lexer.IsSym(s.At())){
					sym+=s.At()
					s.Next()
				}
				
				s.BuildToken(sym, lexer.TKNS(s.At()))
			}

			if (lexer.IsAlpha(s.At())){
				str:=""
				for (lexer.IsAlphaNum(s.At())){
					str+=s.At()
					s.Next()
				}
				
				if (lexer.TKNS(str) == kind.FKTKN){
					s.BuildToken(str, kind.Identifier)
				}else{
					s.BuildToken(str, lexer.TKNS(str))
				}
			}
			
			if (lexer.SKP(s.At())){
				s.Next()
			}
		}
	}

	s.BuildToken("EOF", kind.EOF)

	return s.Tokens

}

func BuildLexer(src string) []tokens.Token {

	lex:=Lexer{Chars:fmt.Strings.Listify(src+" ")}
	return lex.Tokenize()

}