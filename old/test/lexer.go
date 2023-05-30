package test

import (
	"Soup/old/lexer"
	"fmt"
)

func Lexer() {

	for _, v := range lexer.Tokenize("{`test`} -? hello ?- def") {
		fmt.Printf("\nValue: %v\nType: %v\nLocation: \n Line: %v\n Start: %v\n End: %v\n Global: %v\n",
			v.Value,
			v.Type,
			v.Loco.Line,
			v.Loco.Start,
			v.Loco.End,
			v.Loco.Global,
		)
	}

}
