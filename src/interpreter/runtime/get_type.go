package runtime

import (
	f "fmt"
	"os"
)

var Dtype = Func(
	"type",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 1 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 1 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		return Make_String(args[0].GetType())
	},
)
