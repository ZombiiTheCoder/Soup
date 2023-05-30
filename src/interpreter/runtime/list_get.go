package runtime

import (
	f "fmt"
	"os"
)

var DList_Get = Func(
	"list_get",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 1 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 2 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if (args[0].GetType() != "ArrayVal"){
			f.Println("Argument must be List/Array")
			os.Exit(1)
		}

		switch args[1].GetType() {
		case "NumeralVal":
			return args[0].(ArrayVal).Elements[args[1].(NumeralVal).Val]
		default:
			f.Println("Cannot Use Value As Index")
			os.Exit(1)
		}
		return Make_Null()
	},
)
