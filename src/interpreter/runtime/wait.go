package runtime

import (
	f "fmt"
	"os"
	"time"
)

var DWait = Func(
	"wait",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		if len(args) > 1 {
			f.Printf("\nToo Many Values ( %v ) ", len(args))
			os.Exit(1)
		}

		if len(args) < 1 {
			f.Printf("\nToo Little Values ( %v ) ", len(args))
			os.Exit(1)
		}

		wait := time.Duration(GetVal(args[0]).(int)) * time.Second
		if args[0].GetType() == "NumeralVal" {
			time.Sleep(wait)
		} else {
			f.Printf("\nType ( %v ) Cannot Be Used For Wait Func", args[0].GetType())
			os.Exit(1)
		}
		return Make_Null()
	},
)
