package arrays

type MapObject struct {
	Key   any
	Value any
}

type Map struct {
	Obj  any
	Next any
}

func NewMap() Map {
	return Map{}
}