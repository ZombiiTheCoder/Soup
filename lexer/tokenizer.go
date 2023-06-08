package lexer

import (
	"regexp"
	"soup/tokens"
	"soup/utils"
)

func (s *Lexer) Tokenize() []tokens.Token {

	for s.I < len(s.Chars) {

		s.LexOCT()
		s.InlineComments()
		s.BlockComments()
		s.LexString()
		s.LexNumber()
		s.LexSymbol()
		s.LexAlpha()

		s.SkipChar()
		IsValid := func () bool { match, _ := regexp.MatchString(`^([0-9]|[a-zA-Z_]|[=!&~><+*%\/.?]|[-]|[|]|["'`+"`"+`])*$`, s.At()); return match }
		if !IsValid() && s.At() != ";" && s.At() != ":" && s.At() != "," && s.At() != "\n" && s.At() != " " && s.At() != "\t" && s.At() != "\r" && s.At() != "\n\r" && s.At() != "(" && s.At() != ")" && s.At() != "{" && s.At() != "}" && s.At() != "[" && s.At() != "]" {
			utils.Error(
				"Invalid Char %v on Line: %v Column: %v\n %v:%v:%v",
				s.At(),
				s.Line,
				s.Column,
				s.FileName,
				s.Line,
				s.Column,
			)
		}
		
	}

	s.AddToken(tokens.ConstructToken("EOF", tokens.EndOfFile, s.FileName, s.Line, s.Column))

	return s.Tokens

}