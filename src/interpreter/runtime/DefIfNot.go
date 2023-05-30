package runtime

import (
	f "fmt"
	"os"
)

var DEFIFNOT = Func(
	"DefIfNot",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 3 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 3 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if _, ok := scope.Vars[args[0].(StringVal).Val]; !ok {
			scope.DeclareVar(args[0].(StringVal).Val, args[1], args[2].(BooleanVal).Val)
		}

		return Make_Null()
	},
)
