package token

import "Soup/src/lex2/token/kind"

type Location struct {
	Line int
	Start int
	End int
	Global int
}

type Token struct {
	Value string
	Type kind.TokenType
	Literal bool
	Location
}