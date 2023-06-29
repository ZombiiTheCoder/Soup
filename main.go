package main

// import (
// 	"soup/spoon"
// 	"soup/utils"
// )

import (
	"fmt"
	Arrays "soup/lib/arrays"
)

var Solution string

func main(){

	// if (Solution == "soup"){
		// Soup()
	// }else if (Solution == "spoon"){
		// spoon.Spoon()
	// }else {
		// utils.Error("Invalid Solution Type: %v", Solution)
	// }

	nw := Arrays.NewArray()
	
	for i := 0; i < 101; i++ {
		nw.Add("Hello World"+fmt.Sprintf("%v", i))
	}

	nw.Set(100, "Working Add, Set, and Find features in Arrays :)")

	for i := 0; i < 101; i++ {
		fmt.Println(nw.Find(i))
	}
	

}