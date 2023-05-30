package kind

type TokenKind int

const (
	FKTKN TokenKind = iota
	EOF

	OpenParen
	ClosedParen

	OpenBrace
	ClosedBrace

	OpenBracket
	ClosedBracket

	GreaterThan
	LessThan
	LTEquals
	GTEquals

	Ampersan
	Star
	Exclamation

	Semicolon
	Colon
	Percent
	Slash
	Period
	Comma

	Equals
	DEquals
	DNEquals
	TEquals
	TNEquals

	Plus
	Minus
	DPlus
	DMinus
	Tilde

	Comments

	// Keywords
	Def
	Var
	Soup
	If
	Else
	Elif
	While

	// Types
	Numeral
	Identifier
	String
)
