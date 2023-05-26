package fmt

import (
	f "fmt"
	"os"
)

func (s *Prt) Print(args... any){

	for _, v := range args {
		print(v.(string))
	}

}

func (s *Prt) PrintLn(args... any){

	for _, v := range args {
		print(v.(string))
	}
	print("\n")

}

func (s *Prt) PrintInd(args... any){
	
	for _, v := range args {
		print(v.(string))
		print("\n")
	}

}

func (s *Prt) Error(str string) {

	s.PrintLn(str)
	os.Exit(1)

}

func (s *Prt) ErrorF(str string, args... any) {

	f.Printf(str, args)
	os.Exit(1)

}