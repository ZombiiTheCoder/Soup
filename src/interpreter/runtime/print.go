package runtime

import f "fmt"

var DPrint = Func(
	"Native_Print",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		for _, v := range args {
			f.Print(f.Sprint(GetVal(v)) + " ")
		}

		return Make_Null()
	},
)

var DPrintLn = Func(
	"Native_Println",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		for _, v := range args {
			f.Print(f.Sprint(GetVal(v)) + " ")
		}
		f.Print("\n")

		return Make_Null()
	},
)
