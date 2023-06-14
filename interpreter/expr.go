package interpreter

import (
	"fmt"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
)

func (s *Interpreter) EvalIdentifier(node ast.Identifier, env runtime.Env) runtime.Val {

	return env.LookUpVar(node.Symb)

}

func (s *Interpreter) EvalObject(node ast.ObjectLiteral, env runtime.Env) runtime.Val {

	obj := runtime.Object{ObjectElements: make(map[string]runtime.Val), Type: "Object"}
	for _, v := range node.Properties {
		key := v.Key
		val := v.Val.(ast.Expr)
		var newVal runtime.Val
		if val == nil {
			newVal = env.LookUpVar(key)
		}else {
			newVal = s.Eval(val, env)
		}

		obj.ObjectElements[key] = newVal
	}
	
	return obj

}

func (s *Interpreter) EvalArray(node ast.ArrayExpr, env runtime.Env) runtime.Val {

	Elements:=make([]runtime.Val, 0)
	for _, v := range node.Elements {
		Elements = append(Elements, v)
	}

	return runtime.Array{Type: "Array", Elements: Elements}

}

func (s *Interpreter) EvalCall(node ast.CallExpr, env runtime.Env) runtime.Val {
	args := make([]runtime.Val, 0)

	for _, v := range node.Args {
		args = append(args, s.Eval(v, env))
	}
	function := s.Eval(node.Caller, env)

	if function.GetType() != "NativeFunc" && function.GetType() != "Func" {
		
		utils.Error(`Called Type %v That Is Not Of Function` , function.GetType())

	}

	if function.GetType() == "NativeFunc" {
		result := function.(runtime.NativeFunc).Call(args, env)
		return result
	}

	if function.GetType() == "Func" {
		
		functionCallable := function.(runtime.Func)
		scope := runtime.CreateEnvWithParent(functionCallable.DecEnv)

		if len(functionCallable.Params) > len(args) || len(functionCallable.Params) < len(args) {
			utils.Error("To Little/Many Args provided to function %v of %v of name %v", len(functionCallable.Params), len(args), functionCallable.Name)
		}
		i:=0
		for _, v := range functionCallable.Params {
			
			scope.DeclareVar(v, args[i], false)
			i++
		}

		var result runtime.Val = runtime.Null{Type: "Null", Value: "null"}

		for _, v := range functionCallable.Body {
			va := s.Eval(v, scope)
			if va.GetType() == "Return" {
				result = va.(runtime.Return).Val
				break
			}
		}

		return result
	}

	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalMemberTypes(node ast.MemberExpr, env runtime.Env) runtime.Val {

	if node.Computed{
		switch node.GetType() {
		case "MemberExpr":
			return s.EvalMember(node, env)
		case "ArrayExpr":
			return s.EvalArrayObj(node, env)
		case "StringLiteral":
			return s.EvalStringObj(node, env)
		default:
			utils.Error("Type %v Cannot Use Computed Type", node.GetType())
		}
	}else {
		switch node.GetType() {
		case "MemberExpr":
			return s.EvalMember(node, env)
		default:
			utils.Error("Type %v Cannot Use Dot Operator", node.GetType())
		}
	}
	return runtime.Null{Type: "Null", Value: "null"}

}

func (s *Interpreter) EvalMember(node ast.MemberExpr, env runtime.Env) runtime.Val {

	var (
		NObj ast.Expr = node
		NProp any
	)
	Able := true

	keys := make([]string, 0)

	for Able {
		
		switch NObj.GetType() {
		case "MemberExpr":
			Able = true
			if NObj.(ast.MemberExpr).Property != nil && NObj.(ast.MemberExpr).Property.GetType() == "IntegerLiteral" {
				return s.EvalArrayObj(NObj.(ast.MemberExpr), env)
			}
			NProp = NObj.(ast.MemberExpr).Property
			NObj = NObj.(ast.MemberExpr).Obj
		case "Identifier":
			Able = false
		case "StringLiteral":
			Able = false
		default:
			Able = false
		}

		if NProp != nil {
			switch NProp.(ast.Stmt).GetType() {
			case "Identifier":
				keys=append(keys, NProp.(ast.Identifier).Symb)
				
				switch NObj.GetType() {
				case "MemberExpr":
					Able = true
					NProp = NObj.(ast.MemberExpr).Property
				case "Identifier":
					Able = false
				case "StringLiteral":
					Able = false
				default:
					Able = false
				}
			case "StringLiteral":
				keys=append(keys, NProp.(ast.StringLiteral).Valu)
				
				switch NObj.GetType() {
				case "MemberExpr":
					Able = true
					NProp = NObj.(ast.MemberExpr).Property
				case "Identifier":
					Able = false
				case "StringLiteral":
					Able = false
				default:
					Able = false
				}
			}
		}	

	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	parent:=s.Eval(NObj, env)
	for _, v := range keys {

		parent = parent.(runtime.Object).ObjectElements[v]

	}


	fmt.Println(keys)

	return parent

}

func (s *Interpreter) EvalArrayObj(node ast.MemberExpr, env runtime.Env)  runtime.Val {

	index := node.Property.(ast.IntegerLiteral).Valu
	var array ast.Expr
	fmt.Println(node)
	array = s.Eval(node.Obj.(ast.Stmt), env)

	return s.Eval(array.(runtime.Array).Elements[index], env)

}