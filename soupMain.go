package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"soup/interpreter"
	"soup/utils/builder"
	"strings"
)


func Soup() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
	
		default:
			log.Fatalln("Invalid flag")
		case "help":
			fmt.Println("Load Soup's CLI runner")
			fmt.Println(`	soup.exe`)
			fmt.Println("Run {file.soup}")
			fmt.Println(`	soup.exe -r {file.soup}`)
			fmt.Println("Tokenize {file.soup}")
			fmt.Println(`	soup.exe -l {file.soup} {file}.tokens.json`)
			fmt.Println("Parse {file.soup}")
			fmt.Println(`	soup.exe -p {file.soup} {file}.ast.json`)
		case "-r":
			builder.CreateNewIntepreter(os.Args[2], nil)
		case "-l":
			f, _ := os.Create(os.Args[3])
			tkns:=builder.CreateNewLexer(os.Args[2])
			b, _ := json.MarshalIndent(tkns, "", "    ")
			f.WriteString(string(b))
			f.Close()
		case "-p":
			f, _ := os.Create(os.Args[3])
			ast:=builder.CreateNewParser(os.Args[2])
			b, _ := json.MarshalIndent(ast, "", "    ")
			f.WriteString(string(b))
			f.Close()
		}
	os.Exit(0)
	}
	env := interpreter.CreateEnv()
	fmt.Println(
`Soup v0.1.3
------------------------`,
)
	for {
		fmt.Print("soup >> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Print("\n")
		if strings.Contains(strings.TrimSpace(scanner.Text()), "exit") 
		strings.Contains(strings.TrimSpace(scanner.Text()), "quit") {
			os.Exit(1)
		}
		fmt.Println(builder.CreateNewIntepreterWithText(strings.TrimSpace(scanner.Text()), env).GetValue())

	}
}