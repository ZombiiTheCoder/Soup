package interpreter

import (
	"Soup/src/ast"
	rt "Soup/src/interpreter/runtime"
	f "fmt"
	"math"
	"os"
)

func (s *Inte) Eval_binary_expr(node ast.BinaryExpr, env rt.Env) rt.RuntimeVal {

	left := s.Eval(node.Left, env)
	right := s.Eval(node.Right, env)

	if rt.IsNumeric(left) && rt.IsNumeric(right) {
		return s.Eval_numeric_binary_expr(
			left,
			right,
			node.Op,
		)
	} else {
		return s.Eval_nonnumerical_binary_expr(
			left,
			right,
			node.Op,
		)
	}

}

func (s *Inte) Eval_numeric_binary_expr(left, right rt.RuntimeVal, op string) rt.RuntimeVal {

	result := 0.0
	l := 0.0
	r := 0.0
	var out string = "int"

	if left.GetType() == "FloatLiteral" || right.GetType() == "FloatLiteral" {
		result = 0.0
		out = "float"
	}
	if left.GetType() == "FloatLiteral" {
		l = left.(rt.FloatVal).Val
	} else {
		l = float64(left.(rt.NumeralVal).Val)
	}
	if right.GetType() == "FloatLiteral" {
		r = right.(rt.FloatVal).Val
	} else {
		r = float64(right.(rt.NumeralVal).Val)
	}
	switch op {
	case "+":
		result = l + r
	case "-":
		result = l - r
	case "/":
		result = l / r
	case "*":
		result = l * r
	case "%":
		result = math.Mod(float64(l), float64(r))

	}

	if out == "float" {
		return rt.Make_Float(float64(result))
	} else if out == "int" {
		return rt.Make_Numeral(int(result))
	}

	return rt.Make_Numeral(0)

}

func (s *Inte) Eval_nonnumerical_binary_expr(left, right rt.RuntimeVal, op string) rt.RuntimeVal {

	var result string
	var l any
	var r any

	switch left.GetType() {
	case "NumeralVal":
		l = left.(rt.NumeralVal).Val

	case "FloatVal":
		l = left.(rt.NumeralVal).Val

	case "StringVal":
		l = left.(rt.StringVal).Val

	case "BooleanVal":
		l = left.(rt.BooleanVal).Val

	case "NullVal":
		l = "null"
	}

	switch right.GetType() {
	case "NumeralVal":
		r = right.(rt.NumeralVal).Val

	case "FloatVal":
		r = right.(rt.NumeralVal).Val

	case "StringVal":
		r = right.(rt.StringVal).Val

	case "BooleanVal":
		r = right.(rt.BooleanVal).Val

	case "NullVal":
		r = "null"
	}

	switch op {
	case "+":
		result = f.Sprint(l) + f.Sprint(r)
	default:
		f.Printf("Cannot Use %v Operator On Non Int/Float Literals\n", op)
		os.Exit(1)

	}

	return rt.Make_String(result)

}

func (s *Inte) Eval_Unary_expr(ex ast.UnaryExpr, env rt.Env) rt.RuntimeVal {

	ident := ex.Left
	if ex.Left.GetType() != "Identifier" && ex.Op == "++" ||
		ex.Left.GetType() != "Identifier" && ex.Op == "--" {
		f.Printf("Value %v is not identifier\n", ex.Left)
		os.Exit(1)
	}
	if ex.Prefix {
		return s.Eval_Prefixed_Unary_Expr(
			ident,
			ex.Op,
			env,
		)
	} else {
		return s.Eval_Postfixed_Unary_Expr(
			ident,
			ex.Op,
			env,
		)
	}

}

func (s *Inte) Eval_Prefixed_Unary_Expr(Ident ast.Expr, op string, env rt.Env) rt.RuntimeVal {

	lft := s.Eval(Ident.(ast.Stmt), env)
	var result float64 = 0.0
	if lft.GetType() == "NumeralVal" {
		result = float64(lft.(rt.NumeralVal).Val)
	} else {
		result = float64(lft.(rt.FloatVal).Val)
	}

	switch op {
	case "++":
		result += 1
	case "--":
		result -= 1
	case "+":
		result = +result
	case "-":
		result = -result
	case "~":
		result = float64(int(((result * -1) - 1)))
	}

	switch lft.GetType() {
	case "NumeralVal":
		if op == "++" || op == "--" {
			env.AssignVar(Ident.(ast.Identifier).Symb, rt.Make_Numeral(int(result)))
		}
		return rt.Make_Numeral(int(result))
	case "FloatVal":
		if op == "++" || op == "--" {
			env.AssignVar(Ident.(ast.Identifier).Symb, rt.Make_Float(float64(result)))
		}
		return rt.Make_Float(float64(result))
	}

	return rt.Make_Null()

}

func (s *Inte) Eval_Postfixed_Unary_Expr(Ident ast.Expr, op string, env rt.Env) rt.RuntimeVal {

	lft := s.Eval(Ident.(ast.Stmt), env)
	var result float64 = 0.0
	if lft.GetType() == "NumeralVal" {
		result = float64(lft.(rt.NumeralVal).Val)
	} else {
		result = float64(lft.(rt.FloatVal).Val)
	}
	var lst float64 = result
	switch op {
	case "++":
		result += 1
	case "--":
		result -= 1
	case "+":
		result = +result
	case "-":
		result = -result
	case "~":
		result = float64(int(((result * -1) - 1)))
	}

	switch lft.GetType() {
	case "NumeralVal":
		env.AssignVar(Ident.(ast.Identifier).Symb, rt.Make_Numeral(int(result)))
		return rt.Make_Numeral(int(lst))
	case "FloatVal":
		env.AssignVar(Ident.(ast.Identifier).Symb, rt.Make_Float(float64(result)))
		return rt.Make_Float(float64(lst))
	}

	return rt.Make_Null()

}

func (s *Inte) Eval_identifier(Ident ast.Identifier, env rt.Env) rt.RuntimeVal {

	val := env.LookUpVar(Ident.Symb)
	return val

}

func (s *Inte) Eval_Assign_Expr(node ast.AssignExpr, env rt.Env) rt.RuntimeVal {
	if node.Assigner.GetType() != "Identifier" {
		f.Printf("\nInvalid Identifier %v", node.Assigner)
		os.Exit(1)
	}
	varname := node.Assigner.(ast.Identifier).Symb
	return env.AssignVar(varname, s.Eval(node.Valu, env))
}

func (s *Inte) Eval_object_expr(obj ast.ObjectLiteral, env rt.Env) rt.RuntimeVal {

	object := rt.Make_ObjectVal(make(map[string]rt.RuntimeVal)).(rt.ObjectVal)
	for _, q := range obj.Properties {
		k := q.Key
		v := q.Val.(ast.Expr)
		var runtimeVal rt.RuntimeVal
		if v == nil {
			runtimeVal = env.LookUpVar(k)
		} else {
			runtimeVal = s.Eval(v, env)
		}

		object.ObjElements[k] = runtimeVal
	}

	return object

}

func (s *Inte) Eval_member_expr(expression ast.MemberExpr, env rt.Env) rt.RuntimeVal {

	nn := false
	nc := false
	keys := make([]string, 0)
	var expr ast.Expr = expression
	var nest any

	keys = append(keys, expression.Property.(ast.Identifier).Symb)
	for !nn {

		switch expr.(ast.Stmt).GetType() {
		case "MemberExpr":
			nc = false
		case "Identifier":
			nc = true
		}

		if !nc {
			expr = expr.(ast.MemberExpr).Obj
			switch expr.(ast.Stmt).GetType() {
			case "MemberExpr":
				p := expr.(ast.MemberExpr).Property
				keys = append(keys, p.(ast.Identifier).Symb)
			case "Identifier":
				keys = append(keys, expr.(ast.Identifier).Symb)
			}
		} else {
			if expr.GetType() == "NullVal" {
				f.Printf("\nObject Does Not Exist\n")
				os.Exit(1)
			}
			parent := s.Eval(expr, env)
			nest = parent.(rt.ObjectVal)

			for i := 1; i < len(keys); i++ {
				nest = nest.(rt.ObjectVal).ObjElements[keys[(len(keys)-1)-i]]
				if nest == nil {
					f.Printf("\nKey '%v' Does Not Exist On Object\n", keys[(len(keys)-1)-i])
					os.Exit(1)
				}
				switch nest.(rt.RuntimeVal).GetType() {
				case "NullVal", "NumeralVal", "FloatVal", "BooleanVal", "StringVal", "FuncVal", "NativeFuncVal", "ObjectVal":
					nn = true
				}
			}
		}

	}

	parent := s.Eval(expr, env)
	switch parent.GetType() {
	case "Object":
		return parent
	default:
		return nest.(rt.RuntimeVal)
	}

}

func (s *Inte) Eval_call_expr(expr ast.CallExpr, env rt.Env) rt.RuntimeVal {
	args := make([]rt.RuntimeVal, 0)

	for _, v := range expr.Args {
		args = append(args, s.Eval(v, env))
	}
	fn := s.Eval(expr.Caller, env)

	if fn.GetType() == "NativeFuncVal" {
		result := fn.(rt.NativeFuncVal).Call(args, env)
		return result
	}
	if fn.GetType() == "FuncVal" {

		funct := fn.(rt.FuncVal)
		scope := rt.CreateEnvWithParent(funct.DecEnv)

		if len(funct.Params) > len(args) {
			f.Printf("\nTo Little Args Provided To Function %v\n", fn.(rt.FuncVal).Name)
			os.Exit(1)
		}
		if len(funct.Params) < len(args) {
			f.Printf("\nTo Many Args Provided To Function %v\n", fn.(rt.FuncVal).Name)
			os.Exit(1)
		}

		for i := 0; i < len(funct.Params); i++ {
			if len(funct.Params) < 1 {
				break
			}
			if len(args) < 1 {
				break
			}
			scope.DeclareVar(funct.Params[i], args[i], false)
		}

		result := rt.Make_Null()

		for _, v := range funct.Body {
			// if (len(funct.Params) < 1) {break}
			q := s.Eval(v, scope)
			if q.GetType() == "RetVal" {
				result = q.(rt.RetVal).Val
				break
			}
		}

		return result
	}

	f.Printf("\nCannot Call '%v' Because It Does Not Have The Function Identifier\n", fn)
	os.Exit(1)
	return rt.Make_Null()
}

func (s *Inte) Eval_Relational_Expr(expr ast.RelationalExpr, env rt.Env) rt.RuntimeVal {

	Left := s.Eval(expr.Left.(ast.Stmt), env)
	Right := s.Eval(expr.Right.(ast.Stmt), env)
	if rt.IsNumeric(Left) && rt.IsNumeric(Right) {
		return s.Eval_Numeric_Relational_Expr(Left, Right, expr.Op, env)
	} else {
		return s.Eval_NonNumeric_Relational_Expr(Left, Right, expr.Op, env)
	}

}

func (s *Inte) Eval_Numeric_Relational_Expr(Left, Right rt.RuntimeVal, Op string, env rt.Env) rt.RuntimeVal {

	var left float64
	var right float64
	result := false
	if Left.GetType() == "NumeralVal" {
		left = float64(Left.(rt.NumeralVal).Val)
	} else {
		left = float64(Left.(rt.FloatVal).Val)
	}
	if Right.GetType() == "NumeralVal" {
		right = float64(Right.(rt.NumeralVal).Val)
	} else {
		right = float64(Right.(rt.FloatVal).Val)
	}

	switch Op {
	case "==":
		result = (left == right)
	case "!=":
		result = (left != right)
	case ">":
		result = (left > right)
	case ">=":
		result = (left >= right)
	case "<":
		result = (left < right)
	case "<=":
		result = (left <= right)
	}

	return rt.Make_Boolean(result)

}

func (s *Inte) Eval_NonNumeric_Relational_Expr(Left, Right rt.RuntimeVal, Op string, env rt.Env) rt.RuntimeVal {

	result := false
	left := rt.GetVal(Left)
	right := rt.GetVal(Right)

	switch Op {
	case "==":
		result = (left == right)
	case "!=":
		result = (left != right)
	default:
		f.Printf("Cannot Use %v Relational Operator On Non Int/Float Literals\n", Op)
		os.Exit(1)
	}

	return rt.Make_Boolean(result)

}

func (s *Inte) Eval_Array_Expr(array ast.ArrayExpr, env rt.Env) rt.RuntimeVal {

	Elements := make([]rt.RuntimeVal, 0)

	for _, v := range array.Elements {

		Elements = append(Elements, s.Eval(v, env))
		// if v.GetType() == "ArrayExpr" {
		// 	Elements = append(Elements, s.Eval_Array_Expr(v.(ast.ArrayExpr), env))
		// } else {
		// }

	}

	return rt.Make_Array_Val(
		Elements,
	)

}
