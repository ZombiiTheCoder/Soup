package interpreter

import (
	// "Soup/src/utils/fmt"
	f "fmt"
	"Soup/src/parser/ast"
	"math"
	rt "Soup/src/interpreter/runtime"
	. "github.com/ez4o/go-try"
)

func (s *Inte) Eval_binary_expr(node ast.BinaryExpr, env Env) rt.RuntimeVal {

	left := s.Eval(node.Left, env)
	right := s.Eval(node.Right, env)

	if (
		left.GetType() == "NumeralVal" ||
		right.GetType() == "NumeralVal" &&
		left.GetType() == "FloatVal" ||
		right.GetType() == "FloatVal") {
		return s.Eval_numeric_binary_expr(
			left,
			right,
			node.Op,
		)
	}else {
		return s.Eval_nonnumerical_binary_expr(
			left,
			right,
			node.Op,
		)
	}

	return rt.Make_Null()

}

func (s *Inte) Eval_numeric_binary_expr(left, right rt.RuntimeVal, op string) rt.RuntimeVal {

	var result any = 0
	var out string = "int"

	if (left.GetType() == "FloatLiteral" || right.GetType() == "FloatLiteral"){
		result = 0.0
		out = "float"
	}
	switch op {
		case "+":
			if (out == "float"){
				result=left.(rt.FloatVal).Val + right.(rt.FloatVal).Val
			}else if (out == "int"){
				result=left.(rt.NumeralVal).Val + right.(rt.NumeralVal).Val
			}
		case "-":
			if (out == "float"){
				result=left.(rt.FloatVal).Val - right.(rt.FloatVal).Val
			}else if (out == "int"){
				result=left.(rt.NumeralVal).Val - right.(rt.NumeralVal).Val
			}
		case "/":
			if (out == "float"){
				result=left.(rt.FloatVal).Val / right.(rt.FloatVal).Val
			}else if (out == "int"){
				result=left.(rt.NumeralVal).Val / right.(rt.NumeralVal).Val
			}
		case "*":
			if (out == "float"){
				result=left.(rt.FloatVal).Val * right.(rt.FloatVal).Val
			}else if (out == "int"){
				result=left.(rt.NumeralVal).Val * right.(rt.NumeralVal).Val
			}
		case "%":
			result=math.Mod(float64(left.(rt.FloatVal).Val), float64(right.(rt.FloatVal).Val))

	}

	if (out == "float"){
		return rt.Make_Float(result.(float64))
	}else if (out == "int"){
		return rt.Make_Numeral(result.(int))
	}

	return rt.Make_Numeral(0)

}

func (s *Inte) Eval_nonnumerical_binary_expr(left, right rt.RuntimeVal, op string) rt.RuntimeVal {

	var result string
	var l any
	var r any

	switch left.GetType(){
		case "NumeralVal":
			l=left.(rt.NumeralVal).Val
		
		case "FloatVal":
			l=left.(rt.NumeralVal).Val

		case "StringVal":
			l=left.(rt.StringVal).Val

		case "BooleanVal":
			l=left.(rt.BooleanVal).Val

		case "NullVal":
			l="null"
	}

	switch right.GetType(){
	case "NumeralVal":
		r=right.(rt.NumeralVal).Val

	case "FloatVal":
		r=right.(rt.NumeralVal).Val
	
	case "StringVal":
		r=right.(rt.StringVal).Val

	case "BooleanVal":
		r=right.(rt.BooleanVal).Val

	case "NullVal":
		r="null"
}

	switch op {
		case "+":
			result=f.Sprint(l) + f.Sprint(r)

	}

	return rt.Make_String(result)

}

func (s *Inte) Eval_identifier (Ident ast.Identifier, env Env) rt.RuntimeVal {

	val := env.LookUpVar(Ident.Symb)
	return val

}

func (s *Inte) Eval_object_expr (obj ast.ObjectLiteral, env Env) rt.RuntimeVal {

	object := rt.Make_ObjectVal(make(map[string]rt.RuntimeVal)).(rt.ObjectVal)
	for _, q := range obj.Properties {
		k := q.Key
		v := q.Val.(ast.Expr)
		var runtimeVal rt.RuntimeVal
		if (v == nil){ runtimeVal=env.LookUpVar(k)}else {runtimeVal=s.Eval(v, env)}

		object.Val[k] = runtimeVal
	}

	return object

}

func (s *Inte) Eval_member_expr (expression ast.MemberExpr, env Env) rt.RuntimeVal {

	nn := false
	nc := false
	keys := make([]string, 0)
	var expr ast.Expr = expression
	var nest any

	for (!nn){

		f.Println(expr.(ast.Stmt).GetType())

		Try(func (){nc = expr.(ast.MemberExpr).Obj == nil}).Catch(func(){nc=false})
		keys = append(keys, expression.Property.(ast.Identifier).Symb)
		if (!nc){
			expr = expr.(ast.MemberExpr).Obj
			p := expr
//
			keys=append(keys,p.(ast.Identifier).Symb)
		}
		

	}

	parent := s.Eval(expr.(ast.Expr), env)
	switch parent.GetType() {
		case "Object":
			return parent
		default:
			return nest.(rt.RuntimeVal)
	}

}