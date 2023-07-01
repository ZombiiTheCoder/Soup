package lexer

type Types string

const (
	InvalidToken Types = "Invalid"

	OpenParen     Types = "("
	ClosedParen   Types = ")"
	OpenBrace     Types = "{"
	ClosedBrace   Types = "}"
	OpenBracket   Types = "["
	ClosedBracket Types = "]"

	Colon     Types = ":"
	SemiColon Types = ";"
	Comma     Types = ","
	Dot       Types = "."

	Slash        Types = "/"
	Percent      Types = "%"
	QuestionMark Types = "?"

	LessThen    Types = "<"
	LThanEqual  Types = "<="
	GreaterThen Types = ">"
	GThanEqual  Types = ">="

	Or   Types = "||"
	Pipe Types = "|"

	And      Types = "&&"
	Ampersan Types = "&"

	Star   Types = "*"
	Bang   Types = "!"
	Carrot Types = "^"

	NEquals Types = "!="
	DEquals Types = "=="

	Equals       Types = "="
	PlusEquals   Types = "+="
	DashEquals   Types = "-="
	SlashEquals  Types = "/="
	ModuloEquals Types = "%="
	StarEquals   Types = "*="

	Plus  Types = "+"
	DPlus Types = "++"

	Minus  Types = "-"
	DMinus Types = "--"

	If       Types = "If"
	Var      Types = "Var"
	Use      Types = "Import"
	Else     Types = "Else"
	Const    Types = "Const"
	While    Types = "While"
	Return   Types = "Return"
	Function Types = "Function"

	Int        Types = "Int"
	Float      Types = "Float"
	Boolean    Types = "Boolean"
	String     Types = "String"
	Char       Types = "Char"
	Identifier Types = "Identifier"
	Type       Types = "Type"

	EOF Types = "EOF"
)

var validOneCharToken = map[string]bool{
	"___": false,
	"(":   true,
	")":   true,
	"{":   true,
	"}":   true,
	"[":   true,
	"]":   true,
	":":   true,
	";":   true,
	",":   true,
	".":   true,
	"/":   true,
	"%":   true,
	"?":   true,
	"<":   true,
	">":   true,
	"|":   true,
	"&":   true,
	"*":   true,
	"!":   true,
	"^":   true,
	"=":   true,
	"+":   true,
	"-":   true,
	"`":   true,
	"'":   true,
	"\"":  true,
	"\n":  true,
	"\r":  true,
	"\b":  true,
	"\f":  true,
	"\t":  true,
	" ":   true,
}

var Keywords = map[string]Types{
	"____":  InvalidToken,
	"var":   Var,
	"const": Const,

	"use": Use,

	"fn":  Function,
	"ret": Return,

	"if":    If,
	"else":  Else,
	"while": While,

	"Int":    Type,
	"Float":  Type,
	"Char":   Type,
	"String": Type,
	"Bool":   Type,
}

type Token struct {
	Value  string
	Type   Types
	Line   int
	Column int
}

func (s *Lexer) buildToken(Value string, Type Types) Token {
	return Token{Value, Type, s.line, s.column}
}