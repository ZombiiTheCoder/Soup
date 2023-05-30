package runtime

import "fmt"

func FormatArray(elements []RuntimeVal) string {
	list := "["
	for i, v := range elements {

		list += fmt.Sprint(GetVal(v))
		if i != len(elements)-1 {
			list += ", "
		}

	}

	list += "]"

	return list
}

func GetLastObject(elements map[string]RuntimeVal) string {
	l := ""
	for k := range elements {
		l = k
	}
	return l
}

func FormatObject(elements map[string]RuntimeVal, ident int) string {
	obj := "{\n\n"
	last := GetLastObject(elements)
	for k, v := range elements {

		switch v.GetType() {
		case "ObjectVal":
			if k != last {
				obj += (indent(ident+1) + k + " : " + FormatObject(v.(ObjectVal).ObjElements, ident+1) + ",\n\n")
			} else {
				obj += (indent(ident+1) + k + " : " + FormatObject(v.(ObjectVal).ObjElements, ident+1) + "\n\n")
			}
		default:
			if k != last {
				obj += (indent(ident+1) + k + " : " + fmt.Sprint(GetVal(v)) + ",\n")
			} else {
				obj += (indent(ident+1) + k + " : " + fmt.Sprint(GetVal(v)) + "\n")
			}
		}

	}

	if ident == 0 {
		obj += "\n" + indent(ident) + "}\n"
	} else {
		obj += "\n" + indent(ident) + "}"
	}

	return obj
}

func indent(tr int) string {
	q := ""
	for i := 0; i < tr; i++ {
		q += "	"
	}
	return q
}
