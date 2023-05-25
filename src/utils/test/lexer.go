package test

import "fmt"
import "Soup/src/lex2"

func RunLexer ( Code string ) {
	for _, v := range lex2.BuildLexer(Code) {
		fmt.Printf("\nValue: %v\nType: %v\nLocation: \n Line: %v\n Start: %v\n End: %v\n Global: %v\n",
			v.Value,
			v.Type,
			v.FileLocation.Line,
			v.FileLocation.Start,
			v.FileLocation.End,
			v.FileLocation.Global,
		)
	}
}