package ast

import (
	"fmt"
)

func (s Identifier) GetValue() string {
	return s.Symb
}

func (s StringLiteral) GetValue() string {
	return s.Valu
}

func (s NumericLiteral) GetValue() string {
	return fmt.Sprint(s.Valu)
}

func (s NullLiteral) GetValue() string {
	return "null"
}

