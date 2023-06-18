package interpreter

import (
	"math"
	"reflect"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
	"soup/utils/rt"
)

func (s *Interpreter) EvalBinaryTypes(node ast.BinaryExpr, env runtime.Env) runtime.Val {

	left := s.Eval(node.Left, env)
	right := s.Eval(node.Right, env)
	op := node.Op

	if left.GetType() == "Object" ||
		left.GetType() == "Member" ||
		left.GetType() == "NativeFunc" ||
		left.GetType() == "Func" ||
		left.GetType() == "Array" &&
		right.GetType() == "Object" ||
		right.GetType() == "Member" ||
		right.GetType() == "NativeFunc" ||
		right.GetType() == "Func" ||
		right.GetType() == "Array" {
			if op != "!=" && op != "==" {
				utils.Error("Cannot Use %v with %v for Binary Operation %v", left.GetType(), right.GetType(), op)
			}
	}

	switch op {
	case "||", "&&":
		return s.EvalBinaryLogical(left, right, op, env)

	case "|", "&":
		if left.GetType() != "Int" && left.GetType() != "Bool" && right.GetType() != "Int" && right.GetType() != "Bool" {
			utils.Error("Cannot Use Type %v With Type %v in Bitwise Calculations", left.GetType(), right.GetType())
		}
		return s.EvalBinaryBitwise(left, right, op, env)

	case "!=", "==":
		return s.EvalBinaryEquality(left, right, op, env)

	case "<", "<=", ">", ">=":
		if left.GetType() != "Int" && left.GetType() != "Float" && right.GetType() != "Float" && right.GetType() != "Int" {
			utils.Error("Cannot Use Type %v With Type %v in Relation Calculations", left.GetType(), right.GetType())
		}
		return s.EvalBinaryRelational(left, right, op, env)

	case "+", "-":
		if left.GetType() == "Int" || left.GetType() == "Bool" || left.GetType() == "Float" && right.GetType() == "Int" || right.GetType() == "Float" || right.GetType() == "Bool" {
			return s.EvalBinaryAdditive(left, right, op, env)
		} else {
			if op == "+" {
				return s.EvalBinaryConcatenation(left, right, op, env)
			}
			utils.Error("Cannot Use '-' Type %v With Type %v in Plus Calculations", left.GetType(), right.GetType())
		}

		// case "*":
		// case "/":
		// case "%":

	}

	return runtime.Null{Type: "Null", Value: "null"}
}

func (s *Interpreter) EvalBinaryBitwise(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	switch op {
	case "|":
		return runtime.Int{Type: "Int", Value: rt.GetAsInt(left) | rt.GetAsInt(right)}
	case "^":
		return runtime.Int{Type: "Int", Value: rt.GetAsInt(left) ^ rt.GetAsInt(right)}
	case "&":
		return runtime.Int{Type: "Int", Value: rt.GetAsInt(left) & rt.GetAsInt(right)}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalBinaryEquality(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	switch op {
	case "==":
		if left.GetType() == right.GetType() {
			return runtime.Bool{Type: "Bool", Value: reflect.DeepEqual(left, right)}
		}
		return runtime.Bool{Type: "Bool", Value: false}
	case "!=":
		if left.GetType() == right.GetType() {
			return runtime.Bool{Type: "Bool", Value: !reflect.DeepEqual(left, right)}
		}
		return runtime.Bool{Type: "Bool", Value: true}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalBinaryAdditive(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	outputType := "Int"
	if left.GetType() == "Float" || right.GetType() == "Float" {
		outputType = "Float"
	}
	lft := rt.GetAsFloat(left)
	rit := rt.GetAsFloat(right)
	
	switch op + outputType {
	case "+Int":
		return runtime.Int{Type: "Int", Value: int(lft + rit)}
	case "-Int":
		return runtime.Int{Type: "Int", Value: int(lft) - int(rit)}
	case "+Float":
		return runtime.Float{Type: "Float", Value: lft + rit}
	case "-Float":
		return runtime.Float{Type: "Float", Value: lft - rit}
	}
	
	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalBinaryRelational(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	lft := rt.GetAsFloat(left)
	rit := rt.GetAsFloat(right)

	switch op {
	case "<":
		return runtime.Bool{Type: "Bool", Value: lft < rit}
	case ">":
		return runtime.Bool{Type: "Bool", Value: lft > rit}
	case ">=":
		return runtime.Bool{Type: "Bool", Value: lft >= rit}
	case "<=":
		return runtime.Bool{Type: "Bool", Value: lft <= rit}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalBinaryConcatenation(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	return runtime.String{Type: "String", Value: rt.ToString(left) + rt.ToString(right)}

}

func (s *Interpreter) EvalBinaryLogical(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	switch op {
	case "||":
		return runtime.Bool{Type: "Bool", Value: runtime.IsTrue(left) || runtime.IsTrue(right)}
	case "&&":
		return runtime.Bool{Type: "Bool", Value: runtime.IsTrue(left) && runtime.IsTrue(right)}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalBinaryMultiplicative(left runtime.Val, right runtime.Val, op string, env runtime.Env) runtime.Val {

	outputType := "Int"
	if left.GetType() == "Float" || right.GetType() == "Float" {
		outputType = "Float"
	}
	lft := rt.GetAsFloat(left)
	rit := rt.GetAsFloat(right)

	switch op + outputType {
	case "*Int":
		return runtime.Int{Type: "Int", Value: int(lft * rit)}
	case "/Int":
		return runtime.Int{Type: "Int", Value: int(lft / rit)}
	case "%Int":
		return runtime.Int{Type: "Int", Value: int(lft) % int(rit)}
	case "*Float":
		return runtime.Float{Type: "Float", Value: lft * rit}
	case "/Float":
		return runtime.Float{Type: "Float", Value: lft / rit}
	case "%Float":
		return runtime.Float{Type: "Float", Value: math.Mod(lft, rit)}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}