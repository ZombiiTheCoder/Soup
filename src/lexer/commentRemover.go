package lexer

import (
	"strings"
)

func removeSingleLine(contents []string)[]string{
    h := "u3498urt891349f" 
    co := 0;
    next := func () string { if (co+1 != len(contents)) {co++;return contents[co-1];}; return contents[co-1]}
    at := func () string { return contents[co]; }
    src := make([]string, len(contents))
    inString := false
    for (co!=(len(contents))){
        if(at() == "`" && !inString){
            src[co] = at()
            co++
            inString = true
        }else if (at() == "`" && inString){
            inString = false
        }
        if at()+contents[co+1] == "??" && !inString{
            for (at() != "\n" && !inString){
                if (inString) {break}
                next()
                src[co] = h
            }
        }else{
            src[co] = at()
            next()
        }
        if (co == len(contents)-1) {break}

    }

    return strings.Split(strings.ReplaceAll(strings.Join(src, ""), h, ""), "")
}

func removeMultiLine(contents []string)[]string{
    h := "u3498urt891349f" 
    co := 0;
    EOF:=false
    next := func () { if (co+1 != len(contents)) {co++} else {EOF=true}}
    at := func () string { return contents[co]; }
    src := make([]string, len(contents))
    inString := false
    for (co!=len(contents)){
        if (len(contents)-1 == co) {break}
        if (EOF){
            break
        }
        src[co]=contents[co]

        if (at() == "`" && !inString){
            src[co] = at()
            inString = true
        }else if (at() == "`" && inString){
            src[co] = at()
            inString = false
        }
        if at()+contents[co+1] == "-?" && !inString{
            src[co] = h
            src[co+1] = h
            for (!inString){
                if (inString) {break}
                src[co] = h
                next()
                if (at()+contents[co+1] != "?-"){
                    src[co] = h
                }else{
                    break
                }
            }
            src[co] = h
            next()
            src[co] = h
            next()
        }else{
            src[co] = at()
            next()
        }

    }
    
    return strings.Split(strings.ReplaceAll(strings.Join(src, ""), h, ""), "")
}

func RemoveComments(contents string)string{
    qs := strings.Split(contents+"  ", "")
    nsl := removeSingleLine(qs)
    nml := removeMultiLine(nsl)
    return strings.Join(nml, "")
    
}