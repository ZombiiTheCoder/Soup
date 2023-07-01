package main

// import (
// 	"soup/spoon"
// 	"soup/utils"
// )

import (
	"fmt"
	"os"
	"soup/lexer"
)

var Solution string

func main(){

	// if (Solution == "soup"){
		// Soup()
	// }else if (Solution == "spoon"){
		// spoon.Spoon()
	// }else {
		// utils.Error("Invalid Solution Type: %v", Solution)
	// }

	lexer:=lexer.Lexer{}
	f, _:=os.ReadFile("main.soup")
	lexer.Init(string(f))
	tokens:=lexer.Tokenize()
	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}
}