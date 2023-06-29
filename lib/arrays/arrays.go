package array

import (
	Maps "soup/lib/Maps"
)

type ArrayObject struct {
	Key   any
	Value any
}

type Array struct {
	Maps.Map
	Key int
}

func NewArray() Array {
	return Array{}
}