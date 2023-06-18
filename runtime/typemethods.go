package runtime

func (s Null) GetType() string {
	return s.Type
}

func (s Int) GetType() string {
	return s.Type
}

func (s Float) GetType() string {
	return s.Type
}

func (s Bool) GetType() string {
	return s.Type
}

func (s String) GetType() string {
	return s.Type
}

func (s Object) GetType() string {
	return s.Type
}

func (s Member) GetType() string {
	return s.Type
}

func (s NativeFunc) GetType() string {
	return s.Type
}

func (s NativeMethod) GetType() string {
	return s.Type
}

func (s Func) GetType() string {
	return s.Type
}

func (s Return) GetType() string {
	return s.Type
}

func (s Array) GetType() string {
	return s.Type
}