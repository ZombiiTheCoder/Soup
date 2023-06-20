package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"soup/ast"
	"soup/runtime"
	"soup/utils"
	"strings"
)

// reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter text: ")
//     text, _ := reader.ReadString('\n')

func getInput(prompt string, r *bufio.Reader) (string) {
	fmt.Print("\n"+prompt)
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	fmt.Print("\n")
	return strings.TrimSpace(scanner.Text())
}

func Func_TerminalInput(olargs []ast.Expr, args []runtime.Val, scope runtime.Env) runtime.Val {

	
	if len(args) > 1 || len(args) < 1 {
		utils.Error("There must be one argument in input(string) function")
	}
	
	text := getInput(fmt.Sprintf("%v", args[0].GetValue()), bufio.NewReader(os.Stdin))

	type VInterpreter struct {
		Interpreter
		Ast ast.Stmt
		Env runtime.Env
	}

	Interp := VInterpreter{}
	Interp.Env = CreateEnv()
	return Interp.Eval(ast.StringLiteral{Type: "StringLiteral", Valu: text}, Interp.Env)
}