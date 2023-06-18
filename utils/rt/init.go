package rt

import (
	"fmt"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
)

func ToString(value ast.Expr) string {
	switch value.GetType() {
	case "Int":
		return fmt.Sprintf("%v", value.(runtime.Int).Value)
	case "Float":
		return fmt.Sprintf("%v", value.(runtime.Float).Value)
	case "Null":
		return "null"
	case "Bool":
		return fmt.Sprintf("%v", value.(runtime.Bool).Value)
	case "String":
		return value.(runtime.String).Value
	default:
		utils.Error("Invalid Value %v", value)
	}
	return ""
}

func GetAsFloat(value ast.Expr) float64 {
	switch value.GetType() {
	case "Int":
		return float64(value.(runtime.Int).Value)
	case "Float":
		return value.(runtime.Float).Value
	case "Bool":
		if value.(runtime.Bool).Value{
			return 1
		}
		return 0
	default:
		utils.Error("Invalid Value %v", value)
	}
	return 0.0
}

func GetAsInt(value ast.Expr) int {
	switch value.GetType() {
	case "Int":
		return value.(runtime.Int).Value
	case "Bool":
		if value.(runtime.Bool).Value{
			return 1
		}
		return 0
	default:
		utils.Error("Invalid Value %v", value)
	}
	return 0.0
}