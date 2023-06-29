package array

func add(parent Array, i int, value any) Array {
	if parent.Obj == nil {
		parent.Obj = ArrayObject{Key: i, Value: value}
	} else {
		if parent.Next == nil {
			parent.Next = Array{}
			parent.Next = add(parent.Next.(Array), i+1, value)
		} else {
			parent.Next = add(parent.Next.(Array), i+1, value)
		}
	}
	return parent
}

func (s *Array) Add(value any) {

	if s.Obj == nil {
		s.Obj = ArrayObject{Key: 0, Value: value}
	} else {
		if s.Next == nil {
			s.Next = Array{}
			s.Next = add(s.Next.(Array), 1, value)
		} else {
			s.Next = add(s.Next.(Array), 1, value)
		}
	}

}