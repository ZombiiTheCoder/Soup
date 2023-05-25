package parser

import (
	"Soup/src/parser/ast"
	"Soup/src/lex2"
	// "Soup/src/lexer"
	"Soup/src/lex2/token/kind"
	"Soup/src/lex2/token"
	"Soup/src/utils/fmt"
)

type Parser struct {
	
	Ip int
	Tokens []token.Token
	Src string

}

func (this *Parser) At() token.Token{

	return this.Tokens[this.Ip]

}

func (this *Parser) Eat() token.Token{

	this.Ip++
	return this.Tokens[this.Ip-1]

}

func (this *Parser) Expect(ExType kind.TokenType, err string) token.Token{

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

func (this *Parser) ProdAst() ast.Stmt {
	this.Tokens = lex2.BuildLexer(this.Src)
	prg := ast.Create_Program(make([]ast.Stmt, 0)).(ast.Program)

	for this.Not_Eof() {
		prg.Body = append(prg.Body, this.parse_stmt())
	}

	return prg

}

func (this *Parser) parse_stmt() ast.Stmt {

	return this.parse_expr()

}

func (this *Parser) parse_expr() ast.Expr {

	return this.parse_additive_expr()

}

func (this *Parser) parse_additive_expr() ast.Expr {

	left := this.parse_multiplicative_expr()
	for this.At().Value == "+" || this.At().Value == "-" {
		op := this.Eat().Value
		right := this.parse_multiplicative_expr()

		left = ast.Create_BinaryExpr(left, right, op)
	
	}

	return left

}

func (this *Parser) parse_multiplicative_expr() ast.Expr {

	left := this.parse_primary_expr()

	for this.At().Value == "/" || this.At().Value == "*" || this.At().Value == "%" {
	
		op := this.Eat().Value
		right := this.parse_primary_expr()

		left = ast.Create_BinaryExpr(left, right, op)
	
	}

	return left

}

func (this *Parser) parse_primary_expr() ast.Expr {

	tk := this.At().Type

	switch tk {
		case kind.Identifier:
			return ast.Create_Identifier(this.Eat().Value)

		case kind.Numeral:
			return ast.Create_NumericLiteral(this.Eat().Value)
		
		case kind.String:
			return ast.Create_StringLiteral(this.Eat().Value)

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

	return ast.Create_StringLiteral("")

}

func BuildParser(text string) ast.Stmt {
	prse:=Parser{Src: text}
	return prse.ProdAst()
}