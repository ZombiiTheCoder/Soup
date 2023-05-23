package tokens

import "Soup/src/lexer/tokens/kind"
import "Soup/src/lexer/tokens/loco"

type Token struct {
	Value string
	Type kind.TokenKind
	Loco loco.Location
}