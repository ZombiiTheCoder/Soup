package lexer

import (
	"fmt"
	"os"
	"regexp"
	"soup/lib/strings"
)

type Lexer struct {
	pointer int
	chars   []string

	line   int
	column int
	tokens []Token

	eof bool
}

func (s *Lexer) Init(Text string) {
	s.chars = strings.SplitChars(Text)
	s.pointer = 0
	s.line = 1
	s.column = 1
	s.tokens = make([]Token, 0)
	s.eof = false
}

func (s *Lexer) at() string {
	if s.pointer >= len(s.chars) {
		s.eof = true
		return " "
	}else {
		return s.chars[s.pointer]
	}
}

func (s *Lexer) peek(p int) string {
	if s.pointer+p >= len(s.chars) {
		s.eof = true
		return " "
	}else {
		return s.chars[s.pointer+p]
	}
}

func (s *Lexer) next() {
	s.pointer++
}

func (s *Lexer) skippable() {
	switch s.at() {
	case "\n":
		s.line++
		s.column = 0
		s.next()
	case "\r",
	"\b", "\f",
	"\t", "\v",
	"\n\r", " ",
	"":
		s.next()
	}
}

func (s *Lexer) addTkn(value string, Type Types) {
	s.tokens = append(s.tokens, s.buildToken(value, Type))
	s.next()
}

func (s *Lexer) oneCharTokens() {
	switch s.at() {
	case "(": s.addTkn(s.at(), OpenParen)
	case ")": s.addTkn(s.at(), ClosedParen)
	case "{": s.addTkn(s.at(), OpenBrace)
	case "}": s.addTkn(s.at(), ClosedBrace) 
	case "[": s.addTkn(s.at(), OpenBracket)
	case "]": s.addTkn(s.at(), ClosedBracket)
	case ":": s.addTkn(s.at(), Colon)
	case ";": s.addTkn(s.at(), SemiColon)
	case ".": s.addTkn(s.at(), Dot)
	case ",": s.addTkn(s.at(), Comma)
	case "?": s.addTkn(s.at(), QuestionMark)
	case "^": s.addTkn(s.at(), Carrot)
	case "=":
		if s.peek(1) == "=" {
			s.addTkn("==", DEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Equals)
		}
	case "/": 
		if s.peek(1) == "=" {
			s.addTkn("/=", SlashEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Slash)
		}
	case "%":
		if s.peek(1) == "=" {
			s.addTkn("%=", ModuloEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Percent)
		}
	case "+":
		if s.peek(1) == "+" {
			s.addTkn("++", DPlus)
			s.next()
		} else if s.peek(1) == "=" {
			s.addTkn("+=", PlusEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Plus)
		}
	case "-":
		if s.peek(1) == "-" {
			s.addTkn("--", DMinus)
			s.next()
		}else if s.peek(1) == "-=" {
			s.addTkn("-=", DashEquals)
			s.next()
		} else {
			s.addTkn(s.at(), Minus)
		}
	case "*":
		if s.peek(1) == "=" {
			s.addTkn("*=", StarEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Star)
		}
	case ">":
		if s.peek(1) == "=" {
			s.addTkn(">=", GThanEqual)
			s.next()
		}else {
			s.addTkn(s.at(), GreaterThen)
		}
	case "<":
		if s.peek(1) == "=" {
			s.addTkn("<=", LThanEqual)
			s.next()
		}else {
			s.addTkn(s.at(), LessThen)
		}
	case "!":
		if s.peek(1) == "=" {
			s.addTkn("!=", NEquals)
			s.next()
		}else {
			s.addTkn(s.at(), Bang)
		}
	case "&":
		if s.peek(1) == "&" {
			s.addTkn("&&", And)
			s.next()
		}else {
			s.addTkn(s.at(), Ampersan)
		}
	case "|":
		if s.peek(1) == "|" {
			s.addTkn("||", Or)
			s.next()
		}else {
			s.addTkn(s.at(), Pipe)
		}
	}
}

func (s *Lexer) strings() {
	Value := ""
	if s.at() == "\"" {
		s.next()
		for s.at() != "\"" {
			Value = Value + s.at()
			s.next()
		}
		s.addTkn(Value, String)
	}
	if s.at() == "`" {
		s.next()
		for s.at() != "`" {
			Value = Value + s.at()
			s.next()
		}
		s.addTkn(Value, String)
	}
}

func (s *Lexer) char() {
	Value := ""
	if s.at() == "'" {
		s.next()
		for s.at() != "'" {
			Value = Value + s.at()
			s.next()
		}
		if len(Value) > 1 {
			fmt.Println("Invalid Char Value:", Value, "Line:", s.line, "Pos:", s.column)
			os.Exit(1)
		}else {
			s.addTkn(Value, Char)
		}
	}
}

func regx(test, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(test)
}

func (s *Lexer) AlnumTokens() {
	Value := ""
	if regx(s.at(), "[[:alpha:]]") {
		for regx(s.at(), "[[:alnum:]]") {
			Value = Value + s.at()
			s.next()
		}
		if Keywords[Value] == InvalidToken {
			s.tokens = append(s.tokens, s.buildToken(Value, Identifier))
		}else {
			s.tokens = append(s.tokens, s.buildToken(Value, Keywords[Value]))
		}
	}
}

func (s *Lexer) NumTokens() {
	Value := ""
	if regx(s.at(), "^[0-9]*$") {
		for regx(s.at(), "^[0-9]*$") || s.at() == "." {
			Value = Value + s.at()
			s.next()
		}
		if strings.Count(Value, '.') == 0 {
			s.tokens = append(s.tokens, s.buildToken(Value, Int))
		}else if strings.Count(Value, '.') == 1 && regx(Value, "^([0-9.]|[.])+[0-9]*$"){
			s.tokens = append(s.tokens, s.buildToken(Value, Float))	
		}else {
			fmt.Println("Invalid Float Or Int Value:", Value, "Line:", s.line, "Pos:", s.column)
			os.Exit(1)
		}
	}
}


func (s *Lexer) Tokenize() []Token {
	for i := 0; i < len(s.chars); i++ {
		if !s.eof {
			s.column++
			s.skippable()
			s.oneCharTokens()
			s.strings()
			s.char()
			s.AlnumTokens()
			s.NumTokens()
			if !regx(s.at(), "[[:alnum:]]") && !validOneCharToken[s.at()] {
				fmt.Println("Invalid Character:", s.at(), "Line:", s.line, "Pos:", s.column)
				os.Exit(1)
			}
				
		}
	}
	s.tokens = append(s.tokens, s.buildToken("End Of File", EOF))
	return s.tokens
}