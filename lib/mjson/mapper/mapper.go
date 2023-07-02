package mapper

import (
	"fmt"
	"os"
	"soup/lib/mjson/parser"
)

type Mapper struct {}

func (s *Mapper) Eval(node parser.Stmt) any {

	switch node.GetType() {
	case "Program":
		body := make(map[string]any)
		body["body"] = s.Eval(node.(parser.Program).Body)
		return body

	case "Object":
		mapp := make(map[string]any, 0)

		for _, v := range node.(parser.Object).Properties {
			mapp[v.Key] = s.Eval(v.Value)
		}

		return mapp

	case "String":
		return node.(parser.String).Value
	
	case "Int":
		return node.(parser.Int).Value

	case "Float":
		return node.(parser.Float).Value

	case "Null":
		return node.(parser.Null).Value
	
	case "Boolean":
		return node.(parser.Boolean).Value
	
	case "Array":
		arr := make([]any, 0)

		for _, v := range node.(parser.Array).Elements {
			arr = append(arr, s.Eval(v))
		}

		return arr

	default:
		fmt.Println("Unexpected Ast Block of type:", node.GetType())
		os.Exit(1)
	}

	return nil
}