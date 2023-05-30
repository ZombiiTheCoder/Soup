package runtime

func (s NullVal) GetType() string {
	return "NullVal"
}

func (s NumeralVal) GetType() string {
	return "NumeralVal"
}

func (s FloatVal) GetType() string {
	return "FloatVal"
}

func (s BooleanVal) GetType() string {
	return "BooleanVal"
}

func (s StringVal) GetType() string {
	return "StringVal"
}

func (s ObjectVal) GetType() string {
	return "ObjectVal"
}

func (s MemberVal) GetType() string {
	return "MemberVal"
}

func (s NativeFuncVal) GetType() string {
	return "NativeFuncVal"
}

func (s FuncVal) GetType() string {
	return "FuncVal"
}

func (s RetVal) GetType() string {
	return "RetVal"
}

func (s ArrayVal) GetType() string {
	return "ArrayVal"
}
