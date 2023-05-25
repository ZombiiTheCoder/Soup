package lex2

import (
	"Soup/src/utils/fmt"
	"Soup/src/lex2/token"
	"Soup/src/lex2/token/kind"
)

type Lex interface{
	At() string
	Next() int
	BuildToken() int
	Tokenize() []token.Token
}

type Lexer struct{
	Lex
	Chars []string
	Ip int
	Tokens []token.Token
}

func (s *Lexer) At() string {
	return s.Chars[s.Ip]
}

func (s *Lexer) Next() int {
	s.Ip = s.Ip + 1
	return 0
}

func (s *Lexer) BuildToken(value string, typee kind.TokenType) int {

	s.Tokens = append(s.Tokens, token.Token{Value: value, Type: typee})
	return 0

}

func (s *Lexer) Tokenize() []token.Token {

	for (s.Ip < len(s.Chars)){
		
		if (token.IsOneCharToken(s.At())){
			s.BuildToken(s.At(), token.GetTokenType(s.At()))
			s.Next()
		}else{

			if (s.At()+s.Chars[s.Ip+1] == "??"){
				num:=""
				for (s.At() != "\n"){
					num+=s.At()
					s.Next()
				}
			}

			if (s.At()+s.Chars[s.Ip+1] == "-?"){
				num:=""
				for (s.At()+s.Chars[s.Ip+1] != "?-"){
					num+=s.At()
					s.Next()
				}
			}
		
			if (token.IsNumber(s.At())){
				num:=""
				for (token.IsNumber(s.At())){
					num+=s.At()
					s.Next()
				}
				
				s.BuildToken(num, kind.Numeral)
			}
			
			if (token.IsSymbol(s.At())){
				sym:=""
				for (token.IsSymbol(s.At())){
					sym+=s.At()
					s.Next()
				}
				
				s.BuildToken(sym, token.GetTokenType(s.At()))
			}

			if (token.IsAlpha(s.At())){
				str:=""
				for (token.IsAlphaNum(s.At())){
					str+=s.At()
					s.Next()
				}
				
				if (token.GetTokenType(str) == kind.FKTKN){
					s.BuildToken(str, kind.Identifier)
				}else{
					s.BuildToken(str, token.GetTokenType(str))
				}
			}
			
			if (token.IsSkippable(s.At())){
				s.Next()
			}
		}
	}

	s.BuildToken("EOF", kind.EOF)

	return s.Tokens

}

func BuildLexer(src string) []token.Token {

	lex:=Lexer{Chars:fmt.Strings.Listify(src+" ")}
	return lex.Tokenize()

}