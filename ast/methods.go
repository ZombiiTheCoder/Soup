package ast

// Main

func (s BlockStmt) GetType() string {
	return s.Type
}

func (s Program) GetType() string {
	return s.Type
}

// Expressions

func (s AssignExpr) GetType() string {
	return s.Type
}

func (s TernaryExpr) GetType() string {
	return s.Type
}

func (s LogicalExpr) GetType() string {
	return s.Type
}

func (s BinaryExpr) GetType() string {
	return s.Type
}

func (s UnaryExpr) GetType() string {
	return s.Type
}

func (s CallExpr) GetType() string {
	return s.Type
}

func (s MemberExpr) GetType() string {
	return s.Type
}

func (s ArrayExpr) GetType() string {
	return s.Type
}

func (s Property) GetType() string {
	return s.Type
}

// Literals

func (s Identifier) GetType() string {
	return s.Type
}

func (s StringLiteral) GetType() string {
	return s.Type
}

func (s IntegerLiteral) GetType() string {
	return s.Type
}

func (s FloatLiteral) GetType() string {
	return s.Type
}

func (s NullLiteral) GetType() string {
	return s.Type
}

func (s ObjectLiteral) GetType() string {
	return s.Type
}

// Statements

func (s VarDec) GetType() string {
	return s.Type
}

func (s FuncDec) GetType() string {
	return s.Type
}

func (s ImpStmt) GetType() string {
	return s.Type
}

func (s IfStmt) GetType() string {
	return s.Type
}

func (s WhileStmt) GetType() string {
	return s.Type
}

func (s ReturnStmt) GetType() string {
	return s.Type
}
