package main

import (
	"Soup/src/interpreter/runtime"
	"Soup/src/interpreter"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"path"
)

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	  return scanner.Text()
	}

	return ""
}

func readFile() (string) {
	body, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
	return string(body)
}


func main(){
	fmt.Println("Soup v0.0.3")
	fmt.Println("---------------------") 
	env := runtime.CreateEnv()
	qd, _ := filepath.Abs("./slib/nil.soup")
	StdPath := filepath.Dir(qd)
	ex, err := os.Executable()
    ExeLocation := path.Dir(ex)
    if err != nil { log.Fatal(err) }
	if (len(os.Args) > 1){
		contents := readFile()
		q, _ := filepath.Abs(os.Args[1])
		Filepath := filepath.Dir(q)
		interpreter.BuildInterpreter(ExeLocation, StdPath, Filepath, contents, env)
	}else{
		q, _ := filepath.Abs("./slib/nil.soup")
		Filepath := filepath.Dir(q)
		for (true){
			fmt.Print("Soup > ")
			text := scanner()
			_, a := interpreter.BuildInterpreter(ExeLocation, StdPath, Filepath, text, env)
			env = a
			// fmt.Println(runtime.GetVal(e))
		}
	}
}