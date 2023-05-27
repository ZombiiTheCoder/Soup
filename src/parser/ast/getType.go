package ast

func (s Program) GetType() string {
	return "Program"
}

func (s VarDec) GetType() string {
	return "VarDec"
}

func (s FuncDec) GetType() string {
	return "FuncDec"
}

func (s RetStmt) GetType() string {
	return "RetStmt"
}

func (s ImpStmt) GetType() string {
	return "ImpStmt"
}

func (s IfStmt) GetType() string {
	return "IfStmt"
}

func (s WhileStmt) GetType() string {
	return "WhileStmt"
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

func (s FloatLiteral) GetType() string {
	return "FloatLiteral"
}

func (s NullLiteral) GetType() string {
	return "NullLiteral"
}

