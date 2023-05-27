package interpreter

import (
	f "fmt"
	"Soup/src/parser/ast"
	"os"
	"math"
	rt "Soup/src/interpreter/runtime"
)

func (s *Inte) Eval_binary_expr(node ast.BinaryExpr, env rt.Env) rt.RuntimeVal {

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

func (s *Inte) Eval_identifier (Ident ast.Identifier, env rt.Env) rt.RuntimeVal {

	val := env.LookUpVar(Ident.Symb)
	return val

}

func (s *Inte) Eval_Assign_Expr (node ast.AssignExpr, env rt.Env) rt.RuntimeVal {
	if (node.Assigner.GetType() != "Identifier"){
		f.Printf("\nInvalid Identifier %v", node.Assigner)
		os.Exit(1)
	}
	varname := node.Assigner.(ast.Identifier).Symb
	return env.AssignVar(varname, s.Eval(node.Valu, env))
}

func (s *Inte) Eval_object_expr (obj ast.ObjectLiteral, env rt.Env) rt.RuntimeVal {

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

func (s *Inte) Eval_member_expr (expression ast.MemberExpr, env rt.Env) rt.RuntimeVal {

	nn := false
	nc := false
	keys := make([]string, 0)
	var expr ast.Expr = expression
	var nest any
	
	keys = append(keys, expression.Property.(ast.Identifier).Symb)
	for (!nn){

		switch expr.(ast.Stmt).GetType() {
			case "MemberExpr":
				nc = false
			case "Identifier":
				nc=true
		}

		if (!nc){
			expr = expr.(ast.MemberExpr).Obj
			switch expr.(ast.Stmt).GetType() {
				case "MemberExpr":
					p := expr.(ast.MemberExpr).Property
					keys=append(keys,p.(ast.Identifier).Symb)
				case "Identifier":
					keys=append(keys,expr.(ast.Identifier).Symb)
			}
		}else{
			parent := s.Eval(expr, env)
			nest = parent.(rt.ObjectVal)

			for i := 1; i < len(keys); i++ {
				nest = nest.(rt.ObjectVal).Val[keys[(len(keys)-1)-i]]
				if (nest == nil) {
					f.Printf("\nKey '%v' Does Not Exist On Object\n", keys[(len(keys)-1)-i])
					os.Exit(1)
				}
				switch nest.(rt.RuntimeVal).GetType() {
					case "NullVal":
					case "NumeralVal":
					case "FloatVal":
					case "BooleanVal":
					case "StringVal":
					case "ObjectVal":
						nn=true
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

func (s *Inte) Eval_call_expr (expr ast.CallExpr, env rt.Env) rt.RuntimeVal {
	args := make([]rt.RuntimeVal, 0)

	for _, v := range expr.Args {
		args = append(args, s.Eval(v, env))
	}
	fn := s.Eval(expr.Caller, env)

	if (fn.GetType() == "NativeFuncVal"){
		result := fn.(rt.NativeFuncVal).Call(args, env)
		return result
	}
	if (fn.GetType() == "FuncVal"){
		
		funct := fn.(rt.FuncVal)
		scope := rt.CreateEnvWithParent(funct.DecEnv)

		if (len(funct.Params) > len(args)){
			f.Printf("\nTo Little Args Provided To Function %v\n", fn.(rt.FuncVal).Name)
			os.Exit(1)
		}
		if (len(funct.Params) < len(args)){
			f.Printf("\nTo Many Args Provided To Function %v\n", fn.(rt.FuncVal).Name)
			os.Exit(1)
		}
		
		for i := 0; i < len(funct.Params); i++ {
			if (len(funct.Params) < 1) {break}
			if (len(args) < 1) {break}
			scope.DeclareVar(funct.Params[i], args[i], false)
		}

		result := rt.Make_Null()

		for _, v := range funct.Body {
			// if (len(funct.Params) < 1) {break}
			q:=s.Eval(v, scope)
			if (q.GetType() == "RetVal"){
				result = q
			}
		}

		return result
	}
	
	f.Printf("\nCannot Call '%v' Because It Does Not Have The Function Identifier\n", fn)
	os.Exit(1)
	return rt.Make_Null()
} 