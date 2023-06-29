package array

import "fmt"

func find(parent Array, key int) any {

	if parent.Obj.(ArrayObject).Key == key {
		return parent.Obj.(ArrayObject).Value.(string) + " " + fmt.Sprintln(key)
	} else {
		if parent.Next == nil {
			return nil
		} else {
			return find(parent.Next.(Array), key)
		}
	}
}

func (s *Array) Find(key int) any {

	if s.Obj.(ArrayObject).Key == key {
		return s.Obj.(ArrayObject).Value
	} else {
		if s.Next == nil {
			return nil
		} else {
			return find(s.Next.(Array), key)
		}
	}

}