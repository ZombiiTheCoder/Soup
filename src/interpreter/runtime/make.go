package runtime

func Make_Null() RuntimeVal {
	return NullVal{ Val: nul }
}

func Make_Numeral(val int) RuntimeVal {
	return NumeralVal{ Val: val }
}

func Make_Float(val float64) RuntimeVal {
	return FloatVal{ Val: val }
}

func Make_Boolean(val bool) RuntimeVal {
	return BooleanVal{ Val: val }
}

func Make_String(val string) RuntimeVal {
	return StringVal{ Val: val }
}