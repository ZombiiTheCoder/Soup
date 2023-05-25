package token

import "regexp"
import "Soup/src/lex2/token/kind"

func IsOneCharToken (char string) bool {

    if (
    char == "(" ||
    char == ")" ||
    char == "[" ||
    char == "]" ||
    char == "{" ||
    char == "}" ||
    char == ";" ||
    char == "," ||
    char == ":" ||
    char == "."){
        return true
    }

    return false

}

func IsSkippable (char string) bool {

    if (
    char == " " ||
    char == "\n" ||
    char == "\t" ||
    char == "\r" ||
    char == "\n\r"){
        return true
    }

    return false

}

func GetTokenType (wrd string) (kind.TokenType) {

    kwrds := map[interface{}]kind.TokenType {
        false: kind.FKTKN,
        "def": kind.Def,
        "mal": kind.Mal,
        "soup": kind.Soup,
        "Soup": kind.Soup,
        "if": kind.If,
        "else": kind.Else,
        "elif": kind.Elif,
        "while": kind.While,
        "=": kind.Equals,
        "==": kind.DEquals,
        "!=": kind.DNEquals,
        "===": kind.TEquals,
        "!==": kind.TNEquals,
        "<": kind.LessThan,
        ">": kind.GreaterThan,
        "<=": kind.LTEquals,
        ">=": kind.GTEquals,
        "*": kind.Star,
        "/": kind.Slash,
        "%": kind.Percent,
        "&": kind.Ampersan,
        "+": kind.Plus,
        "-": kind.Minus,
        "++": kind.DPlus,
        "--": kind.DMinus,
        "!": kind.Exclamation,
        "~": kind.Tilde,
        "(": kind.OpenParen,
        ")": kind.ClosedParen,
        "{": kind.OpenBrace,
        "}": kind.ClosedBrace,
        "[": kind.OpenBracket,
        "]": kind.ClosedBracket,
        ";": kind.Semicolon,
        ":": kind.Colon,
        ".": kind.Period,
    }

    return kwrds[wrd]
}

func IsAlpha(char string) bool {
    match, _ := regexp.MatchString("^[A-Za-z_]*$", char)
    return match
}

func IsAlphaNum(char string) bool {
    match, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", char)
    return match
}

func IsNumber(char string) bool {
    match, _ := regexp.MatchString("^[0-9.]*$", char)
    return match
}

func IsSymbol(char string) bool {
    match, _ := regexp.MatchString(`^[=!&~><+*%-\/]*$`, char)
    return match
}

func IsString(char string) bool {
    match, _ := regexp.MatchString("^[^`]*$", char)
    return match
}

func IsComment(char string) bool {
    match, _ := regexp.MatchString("^[^?]*$", char)
    return match
}