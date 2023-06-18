package interpreter

import (
	"soup/ast"
	"soup/runtime"
	"soup/utils"
	"soup/utils/rt"
	"strings"
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

func (s *Interpreter) NewArrayVals(len int, env runtime.Env) map[string]runtime.Val {
	array_vals := make(map[string]runtime.Val)
	array_vals["length"] = runtime.Int{Type: "Int", Value: len}
	array_vals["append"] = runtime.NativeMethod{Type: "NativeMethod", Name:"append", Call: func(parent runtime.Val, ast_raw []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

			// fmt.Println(args[0])
			// fmt.Println(parent)

			return runtime.Null{Value: "null", Type: "Null"}
		},
	}

	return array_vals
}

func (s *Interpreter) EvalArray(node ast.ArrayExpr, env runtime.Env) runtime.Val {

	Elements:=make([]runtime.Val, 0)
	for _, v := range node.Elements {
		Elements = append(Elements, s.Eval(v, env))
	}

	return runtime.Array{
		Type: "Array",
		Elements: Elements,
		ObjectElements: s.NewArrayVals(len(node.Elements), env),
	}

}

func (s *Interpreter) EvalCall(node ast.CallExpr, env runtime.Env) runtime.Val {
	args := make([]runtime.Val, 0)

	for _, v := range node.Args {
		args = append(args, s.Eval(v, env))
	}
	var function runtime.Val = s.Eval(node.Caller, env)

	if function == nil {
		utils.Error("Null Method Tried To Be Called on %v", node.Caller)
	}

	switch function.GetType() {

		default:
			utils.Error(`Called Type %v That Is Not Of Function` , function.GetType())
			return runtime.Null{Type: "Null", Value: "null"}

		case "NativeFunc":
			return function.(runtime.NativeFunc).Call(node.Args, args, env)
		
		case "NativeMethod":
			// fmt.Println(node)
			return function.(runtime.NativeMethod).Call(s.Eval(node.Caller, env), node.Args, args, env)
		
		case "Func":
			functionCallable := function.(runtime.Func)
			scope := CreateEnvWithParent(functionCallable.DecEnv)
	
			last_param := functionCallable.Params[len(functionCallable.Params)-1]
			switch last_param {
			default:
				if len(functionCallable.Params) > len(args) || len(functionCallable.Params) < len(args) {
					utils.Error("To Little/Many Args provided to function %v of %v of name %v", len(functionCallable.Params), len(args), functionCallable.Name)
				}

				for i := 0; i < len(functionCallable.Params); i++ {
					if functionCallable.Params[i] == "_args" {
						utils.Error("_args must be at the end of arguments list inside of function %v", functionCallable.Name)
					}
					scope.DeclareVar(functionCallable.Params[i], args[i], false)
				}

			case "_args":
			
				if len(functionCallable.Params)-1 > len(args) {
					utils.Error("To Little Args (not including _args/_arguments) provided to function %v of %v of name %v", len(functionCallable.Params), len(args), functionCallable.Name)
				}

				scope.DeclareVar("_args", s.Eval(ast.ArrayExpr{
					Type: "ArrayExpr",
					Elements: func() []ast.Expr {
						_args:=node.Args[len(functionCallable.Params)-1:]
						if len(_args) != 0 {
							return _args
						}
						return make([]ast.Expr, 0)
					}(),
	
				}, env), false)

				for i := 0; i < len(functionCallable.Params)-1; i++ {
					if functionCallable.Params[i] == "_args" {
						utils.Error("_args must be at the end of arguments list inside of function %v", functionCallable.Name)
					}
					scope.DeclareVar(functionCallable.Params[i], args[i], false)
				}
			}
	
			var result runtime.Val = runtime.Null{Type: "Null", Value: "null"}
	
			for _, v := range functionCallable.Body {
				if v.GetType() == "ReturnStmt" {
					result = s.Eval(v.(ast.ReturnStmt).Value, scope)
					break
				}else{
					s.Eval(v, scope)
				}
			}
	
			return result

	}

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
			if NObj.(ast.MemberExpr).Property != nil && NObj.(ast.MemberExpr).Computed && s.Eval(NObj.(ast.MemberExpr).Property, env).GetType() == "Int" {
				return s.EvalSpecialObj(NObj.(ast.MemberExpr), env)
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

		if parent == nil {
			utils.Error("Key Does Not Exist On Object %v", v)
		}

		switch parent.GetType(){
		case "String":
			parent = parent.(runtime.String).ObjectElements[v]
		case "Array":
			parent = parent.(runtime.Array).ObjectElements[v]
		case "Object":
			parent = parent.(runtime.Object).ObjectElements[v]
		}

	}

	return parent

}

func (s *Interpreter) EvalSpecialObj(node ast.MemberExpr, env runtime.Env)  runtime.Val {

	index := s.Eval(node.Property, env).(runtime.Int).Value
	var array ast.Expr = s.Eval(node.Obj.(ast.Stmt), env)
	switch array.GetType() {
	default:
		if index >= len(array.(runtime.Array).Elements) || index < 0 {
			utils.Error("index of %v does not exist in array of length %v", index, len(array.(runtime.Array).Elements))
		}
		return array.(runtime.Array).Elements[index]
	case "String":
		if index >= len(array.(runtime.String).Value) || index < 0 {
			utils.Error("char of index %v does not exist in string of length %v", index, len(array.(runtime.Array).Elements))
		}
		return s.Eval(ast.StringLiteral{Type: "StringLiteral", Valu: strings.Split(array.(runtime.String).Value, "")[index]}, env)
	}


}

func (s *Interpreter) EvalAssign(node ast.AssignExpr, env runtime.Env) runtime.Val {

	if node.Assigner.GetType() != "Identifier" {
		utils.Error("Value is not of Type 'Identifier' %v", node.Assigner)
	}
	name := node.Assigner.(ast.Identifier).Symb
	return env.AssignVar(name, s.Eval(node.Val, env))

}

func (s *Interpreter) EvalUnary(node ast.UnaryExpr, env runtime.Env) runtime.Val {

	if node.Prefix {

		switch node.Op {
		case "++":
			if node.Argument.GetType() != "Identifier"{utils.Error("Not Able To Use Arg %v with op %v",node.Argument, node.Op)}
			arg := s.Eval(node.Argument, env)
			sarg := rt.GetAsFloat(arg)
			outputType := arg.GetType()
			if outputType == "Int" || outputType == "Bool"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Int{Type: "Int", Value: int(1+sarg)})
				return runtime.Int{Type: "Int", Value: int(sarg)}
			}else if outputType == "Float"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Float{Type: "Float", Value: 1+sarg})
				return runtime.Float{Type: "Float", Value: sarg}
			}
		case "--":
			if node.Argument.GetType() != "Identifier"{utils.Error("Not Able To Use Arg %v with op %v",node.Argument, node.Op)}
			arg := s.Eval(node.Argument, env)
			sarg := rt.GetAsFloat(arg)
			outputType := arg.GetType()
			if outputType == "Int" || outputType == "Bool"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Int{Type: "Int", Value: int(sarg-1)})
				return runtime.Int{Type: "Int", Value: int(sarg)}
			}else if outputType == "Float"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Float{Type: "Float", Value: sarg-1})
				return runtime.Float{Type: "Float", Value: sarg}
			}
		}

	}else {
		switch node.Op {
		case "++":
			if node.Argument.GetType() != "Identifier"{utils.Error("Not Able To Use Arg %v with op %v",node.Argument, node.Op)}
			arg := s.Eval(node.Argument, env)
			sarg := rt.GetAsFloat(arg)
			outputType := arg.GetType()
			if outputType == "Int" || outputType == "Bool"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Int{Type: "Int", Value: int(1+sarg)})
				return runtime.Int{Type: "Int", Value: int(sarg)}
			}else if outputType == "Float"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Float{Type: "Float", Value: 1+sarg})
				return runtime.Float{Type: "Float", Value: 1+sarg}
			}
		case "--":
			if node.Argument.GetType() != "Identifier"{utils.Error("Not Able To Use Arg %v with op %v",node.Argument, node.Op)}
			arg := s.Eval(node.Argument, env)
			sarg := rt.GetAsFloat(arg)
			outputType := arg.GetType()
			if outputType == "Int" || outputType == "Bool"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Int{Type: "Int", Value: int(sarg-1)})
				return runtime.Int{Type: "Int", Value: int(sarg-1)}
			}else if outputType == "Float"{
				env.AssignVar(node.Argument.(ast.Identifier).Symb, runtime.Float{Type: "Float", Value: sarg-1})
				return runtime.Float{Type: "Float", Value: sarg-1}
			}
		}
	}

	return runtime.Null{Type: "Null", Value: "null"}

}