package main

import (
	"fmt"
	"soup/ast"
	"soup/utils/builder"
)

func Soup(){
	for _, v := range builder.CreateNewParser("test.soup").(ast.Program).Body {
		fmt.Println(v)
	}
}