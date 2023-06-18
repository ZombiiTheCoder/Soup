package native_functions

import (
	"fmt"
	"soup/ast"
	"soup/runtime"
)

func Func_Print(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	// fmt.Print(args)

	for _, v := range args {
		
		fmt.Print(getVal(v))

	}

	return runtime.Null{Type: "Null", Value: "Null"}

}

func Func_RawPrint(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	// fmt.Print(args)

	for _, v := range args {
		
		fmt.Print(v, " ")

	}

	return runtime.Null{Type: "Null", Value: "Null"}

}