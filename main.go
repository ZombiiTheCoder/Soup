package main

import (
	"soup/spoon"
	"soup/utils"
)

var Solution string

func main(){

	if (Solution == "soup"){
		Soup()
	}else if (Solution == "spoon"){
		spoon.Spoon()
	}else {
		utils.Error("Invalid Solution Type: %v", Solution)
	}

}