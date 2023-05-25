package test

import (
	"Soup/src/parser"
	"fmt"
)

func RunParser ( Code string ) {
	fmt.Println(parser.BuildParser(Code))
}