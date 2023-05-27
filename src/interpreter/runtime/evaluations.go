package runtime

import f "fmt"
// import v8 "github.com/robertkrimen/otto"
import "github.com/traefik/yaegi/interp"
import "github.com/traefik/yaegi/stdlib"
import "os"

var DGoEval = Func(
	"EvalGo",
	func (args []RuntimeVal, scope Env) RuntimeVal{
		
		i := interp.New(interp.Options{})
		i.Use(stdlib.Symbols)
		for _, v := range args {
			if (v.GetType() != "StringVal"){
				f.Printf("\nValue ( %v ) Is Expected To Be Type Of String ", GetVal(v))
				os.Exit(1)
			}
			i.Eval(GetVal(v).(string))
			// if err != nil {
			// 	panic(err)
			// }
		}

		return Make_Null()
	},
)