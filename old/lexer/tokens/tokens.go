package tokens

import (
	"Soup/old/lexer/tokens/kind"
	"Soup/old/lexer/tokens/loco"
)

type Token struct {
	Value string
	Type kind.TokenKind
	Loco loco.Location
}