package lexer

import (
	"soup/tokens"
	"soup/utils"
)

func (s *Lexer) LexString() {
	
	if s.At() == `"` {
		startColumn:=s.Column
		str := ""
		s.Eat()
		for s.At() != `"` {
			str += s.At()
			s.Eat()
			if s.I >= len(s.Chars) {
				utils.Error(
					"Double quote string on line %v column %v Was Not Ended.\n %v:%v:%v",
					s.Line,
					s.Column,
					s.FileName,
					s.Line,
					s.Column,
				)
			}
		}
		s.Eat()
		s.AddToken(tokens.ConstructToken(str, tokens.String, s.FileName, s.Line, startColumn))
	}

	if s.At() == `'` {
		startColumn:=s.Column
		str := ""
		s.Eat()
		for s.At() != `'` {
			str += s.At()
			s.Eat()
			if s.I >= len(s.Chars) {
				utils.Error(
					"Single quote string on line %v column %v Was Not Ended.\n %v:%v:%v",
					s.Line,
					s.Column,
					s.FileName,
					s.Line,
					s.Column,
				)
			}
		}
		s.Eat()
		s.AddToken(tokens.ConstructToken(str, tokens.String, s.FileName, s.Line, startColumn))
	}

	if s.At() == "`" {
		startColumn:=s.Column
		str := ""
		s.Eat()
		for s.At() != "`" {
			str += s.At()
			s.Eat()
			if s.I >= len(s.Chars) {
				utils.Error(
					"Backtick string on line %v column %v Was Not Ended.\n %v:%v:%v",
					s.Line,
					s.Column,
					s.FileName,
					s.Line,
					s.Column,
				)
			}
		}
		s.Eat()
		s.AddToken(tokens.ConstructToken(str, tokens.String, s.FileName, s.Line, startColumn))
	}
	
}