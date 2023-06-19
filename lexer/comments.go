package lexer

import (
	"soup/utils"
)

func (s *Lexer) InlineComments() {
	
	if s.At()+s.Peek() == "??" {
		comment := ""
		s.Eat()
		s.Eat()
		for s.At() != "\n" {
		if s.At() == "\n" { s.Line++; s.Column = 1 }
			comment += s.At()
			s.Eat()
		}
	}

}

func (s *Lexer) BlockComments() {
	
	if s.At()+s.Peek() == "-?" {
		comment := ""
		s.Eat()
		s.Eat()
		for s.At()+s.Peek() != "?-"{
		if s.At() == "\n" { s.Line++; s.Column = 1 }
			comment += s.At()
			s.Eat()
			if s.I >= len(s.Chars) {
				utils.Error(
					"Comment on line %v column %v Was Not Ended.\n %v:%v:%v",
					s.Line,
					s.Column,
					s.FileName,
					s.Line,
					s.Column,
				)
			}
		}
		s.Eat()
		s.Eat()
	}

}