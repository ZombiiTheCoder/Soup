package test

import (
	"Soup/src/lex2"
	"fmt"
)

func RunLexer(Code string) {
	for _, v := range lex2.BuildLexer(Code) {
		fmt.Printf("\nValue: %v\nType: %v\nLocation: \n Line: %v\n Start: %v\n End: %v\n Global: %v\n",
			v.Value,
			v.Type,
			v.Location.Line,
			v.Location.Start,
			v.Location.End,
			v.Location.Global,
		)
	}
}
