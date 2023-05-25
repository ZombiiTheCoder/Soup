package parser

import (
	"Soup/src/lex2"
	// "Soup/src/lexer"
	"Soup/src/lexer/tokens/kind"
	"Soup/src/lexer/tokens"
	"Soup/src/utils/fmt"
)

type Parser struct {
	
	Ip int
	Tokens []tokens.Token
	Src string

}

func (this *Parser) At() tokens.Token{

	return this.Tokens[this.Ip]

}

func (this *Parser) Eat() tokens.Token{

	this.Ip++
	return this.Tokens[this.Ip-1]

}

func (this *Parser) Expect(ExType kind.TokenKind, err string) tokens.Token{

	this.Ip++
	last := this.Tokens[this.Ip-1]
	if (last.Type != ExType){
		fmt.Prints.ErrorF("Parser Error:\n %v - Expecting %v", err, ExType)
	}

	return last

}

func (this *Parser) Not_Eof() bool {

	return this.Tokens[this.Ip].Type != kind.EOF
	
}

func (this *Parser) ProdAst() Stmt {
	this.Tokens = lex2.BuildLexer(this.Src)
	prg := Program{Body: make([]Stmt, 0)}

	for this.Not_Eof() {
		prg.Body = append(prg.Body, this.parse_stmt())
	}

	return prg

}

func (this *Parser) parse_stmt() Stmt {

	return this.parse_expr()

}

func (this *Parser) parse_expr() Expr {

	return this.parse_additive_expr()

}

func (this *Parser) parse_additive_expr() Expr {

	left := this.parse_multiplicative_expr()
	for this.At().Value == "+" || this.At().Value == "-" {
		op := this.Eat().Value
		right := this.parse_multiplicative_expr()

		left = CreateBinaryExpr(left, right, op)
	
	}

	return left

}

func (this *Parser) parse_multiplicative_expr() Expr {

	left := this.parse_primary_expr()

	for this.At().Value == "/" || this.At().Value == "*" || this.At().Value == "%" {
	
		op := this.Eat().Value
		right := this.parse_primary_expr()

		left = CreateBinaryExpr(left, right, op)
	
	}

	return left

}

func (this *Parser) parse_primary_expr() Expr {

	tk := this.At().Type

	switch tk {
		case kind.Identifier:
			return CreateIdentifier(this.Eat().Value)

		case kind.Numeral:
			return CreateNumericLiteral(this.Eat().Value)

		case kind.OpenParen:
			this.Eat()
			value := this.parse_expr()
			this.Expect(
				kind.ClosedParen,
				"Unexpected token found inside parenthesised expression. Expected closing parenthesis.",
			)

			return value

		default:
			fmt.Prints.ErrorF("Unexpected token found during parsing! %v", this.At())
	}

	return CreateStringLiteral("")

}

func BuildParser(text string) Stmt {
	prse:=Parser{Src: text}
	return prse.ProdAst()
}