package kind

type TokenType int64;
const (
    FKTKN TokenType = iota
	EOF 

    OpenParen
    ClosedParen

    OpenBrace
    ClosedBrace

    OpenBracket
    ClosedBracket

    GreaterThan
    LessThan
    LTEquals
    GTEquals

    Ampersan
    Star
    Exclamation
    
    Semicolon
    Colon
    Percent
    Slash
    Period
    Comma

    Equals
    DEquals
    DNEquals
    TEquals
    TNEquals
    
    Plus
    Minus
    DPlus
    DMinus
    Tilde

    Comments

    // Keywords
    Def
	Mal
	Soup
    If
	Else
	Elif
    While
    Season

    // Types
    Numeral
    Identifier
    String
)