package main

import (
	"log"
	"soup/utils/builder"
)

func Soup(){
	// for _, v := range builder.CreateNewParser("test.soup").(ast.Program).Body {
		// fmt.Println(v)
	// }
	log.Fatalln(builder.CreateNewIntepreter("test.soup"))
	// fmt.Println(experimental.Loop(builder.CreateNewParser("test.soup").(ast.Program).Body))
}