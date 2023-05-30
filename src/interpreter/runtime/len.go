package runtime

import (
	f "fmt"
	"os"
)

var DLen = Func(
	"len",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 1 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 1 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		switch args[0].GetType() {
		case "NullVal":
			return nil
		case "NumeralVal":
			f.Println("Cannot Use len() function on Int")
			os.Exit(1)
		case "FloatVal":
			f.Println("Cannot Use len() function on Float")
			os.Exit(1)
		case "BooleanVal":
			f.Println("Cannot Use len() function on Boolean")
			os.Exit(1)
		case "StringVal":
			return Make_Numeral(int(len(args[0].(StringVal).Val)))
		case "ObjectVal":
			return Make_Numeral(int(len(args[0].(ObjectVal).ObjElements)))
		case "MemberVal":
			return Make_Numeral(int(len(args[0].(ObjectVal).ObjElements)))
		case "NativeFuncVal":
			f.Println("Cannot Use len() function on Native Function")
			os.Exit(1)
		case "FuncVal":
			f.Println("Cannot Use len() function on Function")
			os.Exit(1)
		case "ArrayVal":
			return Make_Numeral(int(len(args[0].(ArrayVal).Elements)))
		}
		return Make_Null()
	},
)
