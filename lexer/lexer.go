package lexer

import (
	"regexp"
	"soup/tokens"
	"soup/utils"
	"strings"
)

type Lexer struct {
	I int
	Tokens []tokens.Token
	Chars []string
	FileName string
	Line int
	Column int
}

func (s *Lexer) At() string {
	if s.I >= len(s.Chars){
		return "\n"
	}
	return s.Chars[s.I]
}

func (s *Lexer) Peek() string {
	if s.I+1 < len(s.Chars){
		return s.Chars[s.I+1]
	}
	return " "
}

func (s *Lexer) Eat() {
	s.I++
	s.Column++
}

func (s *Lexer) AddToken(Token tokens.Token){
	s.Tokens = append(s.Tokens, Token)
}

func (s *Lexer) LexOCT() {
	if s.At() == "(" ||
	s.At() == ")" ||
	s.At() == "[" ||
	s.At() == "]" ||
	s.At() == "{" ||
	s.At() == "}" ||
	s.At() == ";" ||
	s.At() == ":" ||
	s.At() == ","{
		s.AddToken(tokens.ConstructToken(s.At(), tokens.GetType(s.At()), s.FileName, s.Line, s.Column))
		s.Eat()
	}
}

func (s *Lexer) SkipChar(){
	if s.At() == "\n" { s.Line++; s.Column = 1 }
	if s.At() == " " ||
	s.At() == "\n" ||
	s.At() == "\f" ||
	s.At() == "\v" ||
	s.At() == "\t" ||
	s.At() == "\r" ||
	s.At() == "\n\r" {
		s.Eat()
	}
	
}

func (s *Lexer) LexNumber() {

	isFloat := func (number string) bool { match, _ := regexp.MatchString("^([0-9.]|[.])+[0-9]*$", number); return match }
	isNumber := func () bool { match, _ := regexp.MatchString("^[0-9]*$", s.At()); return match }
	startColumn:=s.Column
	
	if isNumber() {
		number := ""
		for isNumber() || s.At() == "." {
			number += s.At()
			s.Eat()
		}

		
		if strings.Count(number, ".") == 0 {
			s.AddToken(tokens.ConstructToken(number, tokens.Number, s.FileName, s.Line, startColumn))
		}else if strings.Count(number, ".") == 1 && isFloat(number) {
			s.AddToken(tokens.ConstructToken(number, tokens.Float, s.FileName, s.Line, startColumn))
		}else {
			utils.Error(
				"Invalid Float Or Number %v on line %v column %v Was Not Ended.\n %v:%v:%v",
				number,
				s.Line,
				s.Column,
				s.FileName,
				s.Line,
				s.Column,
			)
		}
	}

}
	

func (s *Lexer) LexSymbol() {

	isSymbol := func () bool { match, _ := regexp.MatchString(`^([=!&~><+*%\/|.?]|[-])*$`, s.At()); return match }
	startColumn:=s.Column

	if isSymbol() {
		combo := ""
		for isSymbol() {
			combo += s.At()
			s.Eat()
		}
		if tokens.GetType(combo) != tokens.InvalidToken{
			s.AddToken(tokens.ConstructToken(combo, tokens.GetType(combo), s.FileName, s.Line, startColumn))
		}else {
			utils.Error(
				"Invalid Symbol Combo %v on line %v column %v Was Not Ended.\n %v:%v:%v",
				combo,
				s.Line,
				s.Column,
				s.FileName,
				s.Line,
				s.Column,
			)
		}
	}

}

func (s *Lexer) LexAlpha() {

	isAlpha := func () bool { match, _ := regexp.MatchString("^[a-zA-Z_]*$", s.At()); return match }
	isAlphaNum := func () bool { match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", s.At()); return match }
	startColumn:=s.Column
	
	if isAlpha() {
		word := ""
		for isAlphaNum() {
			word += s.At()
			s.Eat()
		}
		if tokens.GetType(word) == ""{
			s.AddToken(tokens.ConstructToken(word, tokens.Identifier, s.FileName, s.Line, startColumn))
		}else {
			s.AddToken(tokens.ConstructToken(word, tokens.GetType(word), s.FileName, s.Line, startColumn))
		}
	}

}