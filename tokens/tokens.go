package tokens

type Token struct {
	Value    string
	Type     TokenType
	FileName string
	Line     int
	Column   int
}

func ConstructToken(Value string, Type TokenType, FileName string, Line, Column int) Token {
	// if GetType(){}
	return Token{Value, Type, FileName, Line, Column}
}

func GetType(Value string) TokenType {
	return TokenRecord[Value]
}