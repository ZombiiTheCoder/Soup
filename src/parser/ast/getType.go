package ast

func (s Program) GetType() string {
	return "Program"
}

func (s VarDec) GetType() string {
	return "VarImportDecDec"
}

func (s ImportDec) GetType() string {
	return "ImportDec"
}

func (s AssignExpr) GetType() string {
	return "AssignExpr"
}

func (s Property) GetType() string {
	return "Property"
}

func (s ObjectLiteral) GetType() string {
	return "ObjectLiteral"
}

func (s MemberExpr) GetType() string {
	return "MemberExpr"
}

func (s BinaryExpr) GetType() string {
	return "BinaryExpr"
}

func (s UnaryExpr) GetType() string {
	return "UnaryExpr"
}

func (s CallExpr) GetType() string {
	return "CallExpr"
}

func (s Identifier) GetType() string {
	return "Identifier"
}

func (s StringLiteral) GetType() string {
	return "StringLiteral"
}

func (s NumericLiteral) GetType() string {
	return "NumericLiteral"
}

func (s NullLiteral) GetType() string {
	return "NullLiteral"
}

