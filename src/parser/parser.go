package parser

import (
	"reflect"
	"fmt"
	"os"
	"Soup/src/lexer"
	"Soup/src/lexer/tokens"
	"Soup/src/lexer/tokens/kind"
)

type Parser struct {
	
	Ip int
	Tokenz []tokens.Token

}

func (p *Parser) Next () tokens.Token {

	// if (p.Ip+1 >= len(p.Tokenz)) {
	// 	return tokens.Token { Type: kind.EOF }
	// }
	
	p.Ip++
	prev:=p.Tokenz[p.Ip-1]
	return prev

}


func (p *Parser) Expect (typ kind.TokenKind, err string) tokens.Token {
	
	p.Ip++
	prev:=p.Tokenz[p.Ip-1]
	if (prev.Type != typ){
		fmt.Printf("Parse Error:\n %v %v - Expecting: %v", err, prev, typ)
		os.Exit(1)
	}
	return prev
	
}

func (p *Parser) At () tokens.Token {
	return p.Tokenz[p.Ip];
}

func (p *Parser) Not_Eof () bool {
	if (p.At().Type == kind.EOF){
		return false
	}
	return true
}

func (p *Parser) parse_primay_expr () Expr {

	switch p.At().Type {
		case kind.Identifier:
			return CreateIdentifier(p.Next().Value)
		case kind.Numeral:
			return CreateNumericLiteral(p.Next().Value)
		case kind.String:
			return CreateStringLiteral(p.Next().Value)
		case kind.OpenParen:
			p.Next()
			value := p.parse_expr()
			p.Expect(
				kind.ClosedParen,
				"Non Valid Token Found Between Parens, Need Closing Paren",
			)
			return value
		default:
			fmt.Printf(`Token That Cannot Be Handeled found during Parsing -> %v`, p.At())
			os.Exit(1)
	}

	return CreateStringLiteral(p.Next().Value)

}

func (p *Parser) parse_arguments_list () []Expr {
	args := make([]Expr, 0)
	args = append(args, p.parse_expr())


	for (p.Not_Eof() && p.At().Type == kind.Comma){
		args = append(args, p.parse_assign_expr())
	}

	return args
}

func (p *Parser) parse_arguments () []Expr {

	p.Expect(
		kind.OpenParen,
		"Expected Open Paren For Args List",
	)
	var args []Expr
	if (p.At().Type == kind.ClosedParen){
		args = make([]Expr, 0)
	}else{
		args = p.parse_arguments_list()
	}

	return args

}

func (p *Parser) parse_call_expr (caller Expr) Expr {
	call_expr := CreateCallExpr(caller, p.parse_arguments())

	if (p.At().Type == kind.OpenParen){
		call_expr = p.parse_call_expr(call_expr)
	}

	return call_expr
}

func (p *Parser) parse_member_expr () Expr {

	obj := p.parse_primay_expr()
	var object Expr

	for (p.At().Type == kind.Period || p.At().Type == kind.OpenBracket){

		op := p.Next()
		var property Expr
		var computed bool

		if (op.Type== kind.Period){
			computed = false
			property = p.parse_primay_expr()

			if (reflect.TypeOf(property) != reflect.TypeOf(Identifier{})){
				fmt.Println("Cannot Use Dot Without Right Side Being A Type Of Identifier")
				os.Exit(1)
			}
		
		}else{
			computed = true
			property = p.parse_expr()

			p.Expect(
				kind.ClosedBracket,
				"Missing Closing Bracket In Computed Member Expression",
			)
		
		}

		object = CreateMemberExpr(obj, property, computed)

	}

	return object

}

func (p *Parser) parse_call_member_expr () Expr {
	
	member := p.parse_member_expr()

	if (p.At().Type == kind.OpenParen) {
		
		return p.parse_call_expr(member)
	}

	return member

}

// func (p *Parser) Prefix () tokens.Token {

// 	var op tokens.Token
	
// 	if (
// 		p.At().Type == kind.Plus ||
// 		p.At().Type == kind.Minus ||
// 		p.At().Type == kind.Exclamation ||
// 		p.At().Type == kind.Tilde ||
// 		p.At().Type == kind.DPlus ||
// 		p.At().Type == kind.DMinus){
// 		// token := p.at()
// 		// next()

// 		op = p.At()
// 	}

// 	return op

// }

// func (p *Parser) PostFix () tokens.Token {

// 	var op tokens.Token
	
// 	if (
// 		p.At().Type == kind.DPlus ||
// 		p.At().Type == kind.DMinus){
// 		// token := p.at()
// 		// next()

// 		op = p.At()
// 	}

// 	return op

// }

// func (p *Parser) parse_unary_expression () Expr {
	
	// left := p.parse_call_member_expr()

	// for (
	// 	// p.At().Value == "+" ||
	// 	// p.At().Value == "-" ||
	// 	// p.At().Value == "!" ||
	// 	// p.At().Value == "~" ||
	// 	p.At().Value == "++" ||
	// 	p.At().Value == "--"){
			
	// 	op := p.Next().Value
	// 	right := p.parse_call_member_expr()
	// 	left = CreateBinaryExpr(left, right, op)

	// }parse_unary_expression

	// return left

// }

func (p *Parser) parse_multiplicative_expr () Expr {

	left := p.parse_call_member_expr()

	for (
	p.At().Value == "*" ||
	p.At().Value == "%" ||
	p.At().Value == "/"){

		op := p.Next().Value
		right := p.parse_call_member_expr()
		left = CreateBinaryExpr(left, right, op)

	}

	return left

}

func (p *Parser) parse_additive_expr () Expr {

	left := p.parse_multiplicative_expr()

	for (
	p.At().Value == "+" ||
	p.At().Value == "-"){

		op := p.Next().Value
		right := p.parse_multiplicative_expr()
		left = CreateBinaryExpr(left, right, op)

	}

	return left

}

func (p *Parser) parse_obj_expr () Expr {
	if (p.At().Type != kind.OpenBrace) {
		return p.parse_additive_expr()
	}

	p.Next()
	properties := make([]Property, 0)

	for (p.Not_Eof() && p.At().Type != kind.ClosedBrace){

		key := p.Expect(
			kind.Identifier,
			"Key Expected For Object",
		).Value

		if (p.At().Type == kind.Comma){
			p.Next()
			properties = append(properties, CreateProperty(key, nil))
			continue
		} else if (p.At().Type == kind.ClosedBrace) {
			properties = append(properties, CreateProperty(key, nil))
			continue
		}

		p.Expect(
			kind.Colon,
			"Missing Colon After Key",
		)
		val:= p.parse_expr()

		properties = append(properties, CreateProperty(key, val))
		if (p.At().Type != kind.ClosedBrace){
			p.Expect(
				kind.Comma,
				"Comma Or Closing Brace Required After Value",
			)	
		}
	}

	p.Expect(
		kind.ClosedBrace,
		"Object Missing Closeing Brace",
	)	

	return CreateObjectLiteral(properties)
}

func (p *Parser) parse_assign_expr () Expr {
	left:=p.parse_obj_expr();

	if (p.At().Type == kind.Equals){
		p.Next()
		val:=p.parse_assign_expr()
		return CreateAssignExpr(left, val)
	}

	return left
}

func (p *Parser) parse_expr () Expr {
	return p.parse_assign_expr()
}


func (p *Parser) parse_var_dec () Stmnt {

	isConst := (p.Next().Type == kind.Def)
	ident := p.Expect(
		kind.Identifier,
		"Expected Valid Type For Variable Name",
	).Value
	
	p.Next()

	p.Expect(
		kind.Equals,
		"Expected Equals Symbol To Point Out Var Contents",
	)

	Decl := CreateVarDec(isConst, ident, p.parse_expr())

	return Decl

}

func (p *Parser) parse_stmnt () Stmnt {
	switch p.At().Type {
		case kind.Mal:
		case kind.Def:
			return p.parse_var_dec()

		default:
			return p.parse_expr()
	}
	return p.parse_stmnt()
}

func (p *Parser) ProdAst () Stmnt {
	p.Ip = 0
	prg := Program{
		Body: make([]Stmnt, 0),
	}

	for (p.Not_Eof()){
		// fmt.Println(p.Tokenz[p.Ip], p.Ip)
		prg.Body = append(prg.Body, p.parse_stmnt())
	}

	return prg
}

func CreateParser(text string) Stmnt {
	prse:=Parser{Tokenz: lexer.Tokenize(text)}
	return prse.ProdAst()
}