package lexer

import (
	"fmt"
	"os"
	"strings"

	"Soup/src/lexer/tokens"
	"Soup/src/lexer/tokens/kind"
	"Soup/src/lexer/tokens/loco"
)

func Tokenize(text string) []tokens.Token {
	
	chars := strings.Split(RemoveComments(text), "")
	tokenz := make([]tokens.Token, 0)
	ip := 0
	line := 1
	colum := 1
	EOF := false
	
	
	for ip <= len(chars)-1{
		
		
		BULDTKN := func (VAL string, TKN kind.TokenKind, start, end int) {
			loc := loco.Location{line, start, end, ip}
			ntk := tokens.Token{VAL, TKN, loc}
			tokenz = append(tokenz, ntk)
		}
		
		next := func () { if (ip+1 != len(chars)) {ip++; colum++;}else{EOF=true}}
		at := func () string { return chars[ip]; }
		
		if (EOF){
			break
		}

		// if (at() == "?" && chars[ip+1] == "?"){
		// 	str:=""
		// 	// strt:=colum
		// 	next()
		// 	next()
		// 	for IsComment(at()) && !EOF && at() != "\n" && at() != "\r" && at() != "\n\r"{
		// 		str+=at()
		// 		next()
		// 	}
		// 	// end:=colum+1
		// 	// BULDTKN(str, kind.Comments, strt, end)
		// 	next()
		// 	next()
		// }
		
		// if (at() == "-" && chars[ip+1] == "?"){
		// 	str:=""
		// 	// strt:=colum
		// 	next()
		// 	next()
		// 	for IsComment(at()) && !EOF && at() != "?" && chars[ip+1] != "-"{
		// 		str+=at()
		// 		next()
		// 	}
		// 	// end:=colum+1
		// 	// BULDTKN(str, kind.Comments, strt, end)
		// 	next()
		// 	next()
		// }

		if (OCT(at())){
			BULDTKN(at(), TKNS(at()), colum, colum)
			next()
		}

		if (IsSym(at())){
			tkn:=""
			strt:=colum
			for IsSym(at()) && !EOF{
				tkn+=at()
				next()
			}
			end:=colum
			if (TKNS(tkn) == kind.FKTKN){
				fmt.Printf("Invalid Char Combo '%v' at Line: %v starting at %v and ending at %v", tkn, line, strt, end)
				os.Exit(1)
			}
			BULDTKN(tkn, TKNS(tkn), strt, end)
			// next()
		}

		if (IsAlpha(at())){
			tkn:=""
			strt:=colum
			for IsAlphaNum(at()) && !EOF{
				tkn+=at()
				next()
			}
			end:=colum
			if (TKNS(tkn) == kind.FKTKN){
				BULDTKN(tkn, kind.Identifier, strt, end)
				// fmt.Printf("Invalid Token '%v' at Line: %v starting at %v and ending at %v", tkn, line, strt, end)
				// os.Exit(1)
			}
			BULDTKN(tkn, TKNS(tkn), strt, end)
			next()
		}                                

		if (IsNum(at())){
			tkn:=""
			strt:=colum
			for IsNum(at()) && !EOF{
				tkn+=at()
				next()
			}
			end:=colum
			BULDTKN(tkn, kind.Numeral, strt, end)
			next()
		}
		
		if (at() == "`"){
			str:=""
			strt:=colum
			next()
			for IsString(at()) && !EOF && at() != "`"{
				str+=at()
				next()
			}
			end:=colum+1
			BULDTKN(str, kind.String, strt, end)
			next()
		}

		if (SKP(at())){
			next()
		}

		if (at() == "\n"){
			line++
			colum=0
		}

		if (!IsAlpha(at()) && !IsAlphaNum(at()) && !IsNum("") && !OCT(at()) && !SKP(at()) && !IsSym(at())){
			if (TKNS(at()) == kind.FKTKN){
				fmt.Printf("Invalid Character '%v' at Line: %v starting at %v and ending at %v", at(), line, colum, colum)
				os.Exit(1)
			}
		}

	}

	tokenz = append(tokenz, tokens.Token{ Type: kind.EOF })

	return tokenz
}