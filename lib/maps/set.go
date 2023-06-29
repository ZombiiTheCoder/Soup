package arrays

func set(parent Map, key any, value any) Map {

	if parent.Obj == nil {
		parent.Obj = MapObject{Key: key, Value: value}
	} else {
		if parent.Obj.(MapObject).Key == key {
			parent.Obj = MapObject{Key: key, Value: value}
		} else {
			if parent.Next == nil {
				parent.Next = Map{}
				parent.Next = set(parent.Next.(Map), key, value)
			} else {
				parent.Next = set(parent.Next.(Map), key, value)
			}
		}
	}
	return parent
}

func (s *Map) Set(key any, value any) {

	if s.Obj == nil {
		s.Obj = MapObject{Key: key, Value: value}
	} else {
		if key == "" {
			println("Key Must Not Be An Empty Key:"+key.(string), "Value:", value.(string))
		}
		if s.Obj.(MapObject).Key == key {
			s.Obj = MapObject{Key: key, Value: value}
		} else {
			if s.Next == nil {
				s.Next = Map{}
				s.Next = set(s.Next.(Map), key, value)
			} else {
				s.Next = set(s.Next.(Map), key, value)
			}
		}
	}

}