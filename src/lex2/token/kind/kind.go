package kind

type TokenType int

const (
	FKTKN TokenType = iota
	EOF

	OpenParen
	ClosedParen

	OpenBrace
	ClosedBrace

	OpenBracket
	ClosedBracket

	LessThan
	GTEquals

	LTEquals
	GreaterThan

	Star
	Ampersan
	Exclamation

	Colon
	Slash
	Comma
	Period
	Percent
	Semicolon

	Equals
	DEquals
	DNEquals
	PlusEquals
	MinusEquals
	DivideEquals
	PercentEquals
	MultiplyEquals

	Plus
	Minus
	DPlus
	Tilde
	DMinus

	Comments

	// Keywords
	Fn
	If
	Val
	Var
	Use
	Ret
	Else
	While

	// Types
	Float
	String
	Numeral
	Identifier
)
