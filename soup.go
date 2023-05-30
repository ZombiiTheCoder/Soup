package main

import (
	"Soup/src/interpreter"
	"Soup/src/interpreter/runtime"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func readFile() string {
	body, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func main() {
	env := runtime.CreateEnv()
	qd, _ := filepath.Abs("./pkg/nil.soup")
	StdPath := filepath.Dir(qd)
	ex, err := os.Executable()
	ExeLocation := path.Dir(ex)
	// {
	// 	q, _ := filepath.Abs("./packages/nil.soup")
	// 	Filepath := filepath.Dir(q)
	// 	_, a := interpreter.BuildInterpreter(ExeLocation, StdPath, Filepath, "use `@stdio`", env)
	// 	env = a
	// }
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		contents := readFile()
		q, _ := filepath.Abs(os.Args[1])
		Filepath := filepath.Dir(q)
		interpreter.BuildInterpreter(ExeLocation, StdPath, Filepath, contents, env)
	} else {
		fmt.Println("\nSoup v0.0.4")
		fmt.Println("------------------------")
		q, _ := filepath.Abs("./pkg/nil.soup")
		Filepath := filepath.Dir(q)
		for {
			fmt.Print("Soup > ")
			text := scanner()
			_, a := interpreter.BuildInterpreter(ExeLocation, StdPath, Filepath, text, env)
			env = a
			// fmt.Println(runtime.GetVal(e))
		}
	}
}
