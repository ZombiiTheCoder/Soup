package runtime

import "os"

var DExit = Func(
	"exit",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		os.Exit(0)

		return Make_Null()
	},
)
