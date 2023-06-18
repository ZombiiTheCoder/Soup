package lexer

import (
	"regexp"
	"soup/tokens"
	"soup/utils"
)

func (s *Lexer) LexString() {
	
	if s.At() == `"` {
		startColumn:=s.Column
		str := ``
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
		regx:=regexp.MustCompile(`[!]+[\\]+[nfvtr]`)
		var CHARRCD = map[string]string{
			`791m80v129m023v54821309mv83m94vb509238409mb5m8v23984b52`:`e`,
			`!\n`:"\n",
			`!\f`:"\f",
			`!\v`:"\v",
			`!\t`:"\t",
			`!\r`:"\r",
		}
		var ns = str
		for _, v := range regx.FindAllString(str, -1) {
			ns = regx.ReplaceAllLiteralString(ns, CHARRCD[v])
		}
		s.AddToken(tokens.ConstructToken(ns, tokens.String, s.FileName, s.Line, startColumn))
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