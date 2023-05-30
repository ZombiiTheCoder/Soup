package runtime

import (
	f "fmt"
	"os"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// import v8 "github.com/robertkrimen/otto"

var DGoEval = Func(
	"EvalGo",
	func(args []RuntimeVal, scope Env) RuntimeVal {

		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		var val reflect.Value
		for _, v := range args {
			if v.GetType() != "StringVal" {
				f.Printf("\nValue ( %v ) Is Expected To Be Type Of String ", GetVal(v))
				os.Exit(1)
			}
			val, _ = i.Eval(GetVal(v).(string))
			// if err != nil {
			// 	panic(err)
			// }
		}

		return Make_String(val.String())
	},
)
