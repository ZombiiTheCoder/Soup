package lex2

import (
	"strings"
	"Soup/src/utils/fmt"
	f "fmt"
	"os"
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
	if (s.Ip >= len(s.Chars)){
		return "\n"
	}
	return s.Chars[s.Ip]
}

func (s *Lexer) Peek() string {
	ret := " "
	if (s.Ip+1 < len(s.Chars)){
		ret = s.Chars[s.Ip+1]
	}

	return ret
}

func (s *Lexer) Next() int {
	// if (s.Ip+1 < len(s.Chars)){
		s.Ip++
	// }
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

			// Removes Single Line Comments
			if (s.At()+s.Peek() == "??"){
				num:=""
				for (s.At() != "\n"){
					num+=s.At()
					s.Next()
				}
			}
			
			/* Removes Block Comments */
			if (s.At()+s.Peek() == "-?"){
				num:=""
				s.Next()
				s.Next()
				for (s.At()+s.Peek() != "?-"){
					num+=s.At()
					s.Next()

					if (s.Ip >= len(s.Chars)){
						f.Println("Expected A Combo of ?- For Block Comment")
						os.Exit(1)
					}
				}
				s.Next()
				s.Next()
			}

			/* Checks For String */
			if (s.At() == "`"){
				str:=""
				s.Next()
				for (s.At() != "`"){
					str+=s.At()
					s.Next()

					if (s.Ip >= len(s.Chars)){
						f.Println("Expected ` to Close String")
						os.Exit(1)
					}
				}
				s.Next()

				s.BuildToken(str, kind.String)

			}

			if (token.IsNumber(s.At())){
				num:=""
				for (token.IsNumber(s.At()) || s.At() == "."){
					num+=s.At()
					s.Next()
				}

				if (strings.Count(num, ".") > 1 || !token.ValidateFloat(num)){
					fmt.Prints.ErrorF("Invalid Float  %v", num)
				}else if strings.Count(num, ".") == 0 {
					s.BuildToken(num, kind.Numeral)
				}else if strings.Count(num, ".") == 1 && token.ValidateFloat(num) {
					s.BuildToken(num, kind.Float)
				}
				
			}
			
			if (token.IsSymbol(s.At())){
				sym:=""
				for (token.IsSymbol(s.At())){
					sym+=s.At()
					s.Next()
				}
				if (token.GetTokenType(sym) != kind.FKTKN){
					s.BuildToken(sym, token.GetTokenType(sym))
				}else {
					f.Printf("Invalid Symbol Combo %v\n", sym)
					os.Exit(1)
				}
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

			if (s.At() != "`" && s.At() != ";" && !token.IsSkippable(s.At()) && !token.IsNumber(s.At()) && !token.IsAlpha(s.At()) && !token.IsSymbol(s.At())){
				f.Printf("Invalid Char Found In Source %v\n", s.At())
				os.Exit(1)
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