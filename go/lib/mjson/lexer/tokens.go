package lexer

type Types string

const (
	OpenBrace     Types = "{"
	ClosedBrace   Types = "}"
	OpenBracket   Types = "["
	ClosedBracket Types = "]"
	Comma         Types = ","
	Colon         Types = ":"
	Int           Types = "Int"
	Float         Types = "Float"
	Boolean       Types = "Boolean"
	String        Types = "String"
	Null          Types = "Null"
	EOF           Types = "EOF"
)

type Token struct {
	Value  string
	Type   Types
	Line   int
	Column int
}

func (s *Lexer) buildToken(Value string, Type Types) Token {
	return Token{Value, Type, s.line, s.column}
}