package tokens

type Token struct {
	Value    string    `json:"Value"`
	Type     TokenType `json:"Type"`
	FileName string    `json:"FileName"`
	Line     int       `json:"Line"`
	Column   int       `json:"Column"`
}

func ConstructToken(Value string, Type TokenType, FileName string, Line, Column int) Token {
	// if GetType(){}
	return Token{Value, Type, FileName, Line, Column}
}

func GetType(Value string) TokenType {
	return TokenRecord[Value]
}