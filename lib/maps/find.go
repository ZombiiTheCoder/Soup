package arrays

func find(parent Map, key any) any {

	if parent.Obj.(MapObject).Key == key {
		return parent.Obj.(MapObject).Value
	} else {
		if parent.Next == nil {
			return nil
		} else {
			return find(parent.Next.(Map), key)
		}
	}
}

func (s *Map) Find(key any) any {

	if s.Obj.(MapObject).Key == key {
		return s.Obj.(MapObject).Value
	} else {
		if s.Next == nil {
			return nil
		} else {
			return find(s.Next.(Map), key)
		}
	}

}