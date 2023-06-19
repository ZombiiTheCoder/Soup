package interpreter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
	"strings"
)

// reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter text: ")
//     text, _ := reader.ReadString('\n')

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading input!")
	}
	return strings.TrimSpace(input), err
}

func Func_TerminalInput(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	
	if len(args) > 1 || len(args) < 1 {
		utils.Error("There must be one argument in terminal.input(string) function")
	}
	
	text, _ := getInput(fmt.Sprintf("%v", args[0].GetValue()), bufio.NewReader(os.Stdin))

	fmt.Println(text)

	type VInterpreter struct {
		Interpreter
		Ast ast.Stmt
		Env runtime.Env
	}

	Interp := VInterpreter{}
	Interp.Env = CreateEnv()
	return Interp.Eval(ast.StringLiteral{Type: "StringLiteral", Valu: text}, Interp.Env)
}