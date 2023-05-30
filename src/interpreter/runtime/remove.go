package runtime

import f "fmt"
import "os"

var DRemove = Func(
	"remove",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 1 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 1 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if (args[0].GetType() != "StringVal"){
			f.Println(`Variable Name Must Be String`)
			os.Exit(1)
		}

		if (args[0].(StringVal).Val == "true" ||
			args[0].(StringVal).Val == "false" ||
			args[0].(StringVal).Val == "remove" ||
			args[0].(StringVal).Val == "len" ||
			args[0].(StringVal).Val == "wait" ||
			args[0].(StringVal).Val == "DefIfNot" ||
			args[0].(StringVal).Val == "GoEval"){
			f.Println("Cannot Remove Native Item")
			os.Exit(1)
		}

		if _, ok := scope.Vars[args[0].(StringVal).Val]; ok {
			
			delete(scope.Vars, args[0].(StringVal).Val)
			if _, ok := scope.Consts[args[0].(StringVal).Val]; ok {
				delete(scope.Consts, args[0].(StringVal).Val)
			}
		}

		return Make_Null()
	},
)
