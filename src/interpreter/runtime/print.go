package runtime

import f "fmt"

var DPrint = Func(
	"print",
	func (args []RuntimeVal, scope Env) RuntimeVal{

		for _, v := range args {
			f.Print(f.Sprint(GetVal(v))+" ")
		}
		f.Print("\n")

		return Make_Null()
	},
)