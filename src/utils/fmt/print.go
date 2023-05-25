package fmt

import f "fmt"
import "os"

func (s *Prt) Print(args... any){

	print(args)

}

func (s *Prt) PrintLn(args... any){

	print(args)
	print("\n")

}

func (s *Prt) PrintInd(args... any){
	
	for _, v := range args {
		print(v)
		print("\n")
	}

}

func (s *Prt) Error(str string) {

	s.PrintLn(str)
	os.Exit(1)

}

func (s *Prt) ErrorF(str string, args... any) {

	s.PrintLn(f.Sprintf(str, args))
	os.Exit(1)

}