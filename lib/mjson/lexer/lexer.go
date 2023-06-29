package mjson

import "soup/lib/strings"

type Lexer struct {
	pointer int
	char    byte
	chars   []string

	line   int
	column int
}

func (s *Lexer) Init(Text string) {
	s.chars = strings.SplitChars(Text)
}

func (s *Lexer) Tokenize() {
	for i := 0; i < len(s.chars); i++ {
		
	}
}