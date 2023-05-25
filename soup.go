package main

import "Soup/src/lex2"
// import "Soup/src/lexer"
import "Soup/src/parser"
import "Soup/src/interpreter"
import "fmt"

func main(){
	// for _, v := range lexer.Tokenize("2+5") {
	// 	fmt.Printf("\nValue: %v\nType: %v\nLocation: \n Line: %v\n Start: %v\n End: %v\n Global: %v\n",
	// 	v.Value,
	// 	v.Type,
	// 	v.Loco.Line,
	// 	v.Loco.Start,
	// 	v.Loco.End,
	// 	v.Loco.Global,
	// )
	// }

	for _, v := range lex2.BuildLexer("2+5") {
		fmt.Printf("\nValue: %v\nType: %v\nLocation: \n Line: %v\n Start: %v\n End: %v\n Global: %v\n",
		v.Value,
		v.Type,
		v.Loco.Line,
		v.Loco.Start,
		v.Loco.End,
		v.Loco.Global,
	)
	}

	// fmt.Println(parser.BuildParser("2+2"))
	v := interpreter.Inte{}
	fmt.Println(parser.BuildParser("2+5"))
	fmt.Println(v.Eval(parser.BuildParser("2+5")))
}