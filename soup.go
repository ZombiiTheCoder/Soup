package main

import (
	"Soup/src/interpreter"
	"fmt"
	"bufio"
	"os"
)

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	  return scanner.Text()
	}

	return ""
  }
  

func main(){
	fmt.Println("Soup v0.0.1")
	fmt.Println("---------------------")
	for (true){
		fmt.Print("Soup > ")
		text := scanner()
		fmt.Println(interpreter.BuildInterpreter(text))
	}
}