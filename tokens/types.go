package tokens

type TokenType string

const (
	InvalidToken TokenType = "Invalid_Token"
	EndOfFile              = "EndOfFile"

	OpenParen   = "("
	ClosedParen = ")"

	OpenBrace   = "{"
	ClosedBrace = "}"

	OpenBracket   = "["
	ClosedBracket = "]"

	LessThan = "<"
	LTEquals = "<="

	GTEquals    = ">="
	GreaterThan = ">"

	Or       = "||"
	And      = "&&"
	Star     = "*"
	Bang     = "!"
	Pipe     = "|"
	Carrot   = "^"
	Ampersan = "&"

	Dot          = "."
	Colon        = ":"
	Slash        = "/"
	Comma        = ","
	Percent      = "%"
	Semicolon    = ";"
	QuestionMark = "?"

	NEquals = "!="
	DEquals = "=="

	Equals       = "="
	PlusEquals   = "+="
	DashEquals   = "-="
	SlashEquals  = "/="
	ModuloEquals = "%="
	StarEquals   = "*="

	Plus   = "+"
	Minus  = "-"
	DPlus  = "++"
	Tilde  = "~"
	DMinus = "--"

	// Keywords
	If       = "if"
	Val      = "val"
	Var      = "var"
	Use      = "use"
	Else     = "else"
	While    = "while"
	Return   = "return"
	Function = "function"

	// Types
	Float      = "float"
	String     = "string"
	Number     = "number"
	Identifier = "identifier"
)

var TokenRecord = map[string]TokenType{
	"_____": InvalidToken,

	// Variable
	"val": Val,
	"var": Var,

	// Import
	"use": Use,

	// Function Return Pair
	"fn":  Function,
	"ret": Return,

	// Conditional
	"if":    If,
	"else":  Else,
	"while": While,

	// Equality
	"==": DEquals,
	"!=": NEquals,

	// Relational
	">":  GreaterThan,
	"<":  LessThan,
	">=": GTEquals,
	"<=": LTEquals,

	// Assignment
	"=":  Equals,
	"+=": Equals,
	"-=": Equals,
	"/=": Equals,
	"%=": Equals,
	"*=": Equals,

	// Additive
	"+": Plus,
	"-": Minus,

	// Multiplicative
	"*": Star,
	"/": Slash,
	"%": Percent,

	// Unary
	"!":  Bang,
	"~":  Tilde,
	"++": DPlus,
	"--": DMinus,

	// Access
	".": Dot,
	"(": OpenParen,
	")": ClosedParen,
	"[": OpenBracket,
	"]": ClosedBracket,

	// Bitwise
	"|": Pipe,
	"^": Carrot,
	"&": Ampersan,

	// Logical
	"||": Or,
	"&&": And,

	// Ternary
	":": Colon,
	"?": QuestionMark,

	// Other
	",": Comma,
	"{": OpenBrace,
	"}": ClosedBrace,

	// Semicolon
	";": Semicolon,
}