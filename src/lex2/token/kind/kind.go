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
    Val
	Var
	Fn
    If
	Else
	Elif
    While
    Season
    Use
    Ret

    // Types
    Numeral
    Float
    Identifier
    String
)