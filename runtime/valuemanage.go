package runtime

func (s Null) GetValue() any {
	return s.Value
}

func (s Int) GetValue() any {
	return s.Value
}

func (s Float) GetValue() any {
	return s.Value
}

func (s Bool) GetValue() any {
	return s.Value
}

func (s String) GetValue() any {
	return s.Value
}

func (s Object) GetValue() any {
	return s.ObjectElements
}

func (s Member) GetValue() any {
	return s.ObjectElements
}

func (s NativeFunc) GetValue() any {
	return s.Call
}

func (s NativeMethod) GetValue() any {
	return s.Call
}

func (s Func) GetValue() any {
	return s.Body
}

func (s Return) GetValue() any {
	return s.Value
}

func (s Array) GetValue() any {
	return s.Elements
}