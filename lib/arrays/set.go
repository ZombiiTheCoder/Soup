package array

func set(parent Array, key int, value any) Array {
	if parent.Obj != nil {
		if parent.Obj.(ArrayObject).Key == key {
			parent.Obj = ArrayObject{Key: key, Value: value}
		}
		if parent.Next != nil {
			parent.Next = set(parent.Next.(Array), key, value)
		}
	}
	return parent
}

func (s *Array) Set(key int, value any) {

	if s.Obj != nil {
		if s.Obj.(ArrayObject).Key == key {
			s.Obj = ArrayObject{Key: key, Value: value}
		}
		if s.Next != nil {
			s.Next = set(s.Next.(Array), key, value)
		}
	}

}