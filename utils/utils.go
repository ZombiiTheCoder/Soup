package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Error(err string, args... any){

	fmt.Printf("\n"+err, args...)
	fmt.Print("\n")
	os.Exit(1)

}

func ReadFile(File string) string {
	body, err := ioutil.ReadFile(File)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}