package parser

import (
	"fmt"
	"Soup/src/lexer"
	"Soup/src/lexer/tokens"
	"Soup/src/lexer/tokens/kind"
	"Soup/src/parser/kinds"
)

func Parse(text string){

	ip:=0
	tokenz:=lexer.Tokenize(text)
	next := func () tokens.Token { ip++; prev:=tokenz[ip-1]; return prev}
	expect := func (typ kind.TokenKind, err any) tokens.Token {
		ip++
		prev:=tokenz[ip-1]
		if (prev.Type != typ){
			fmt.Printf("Parse Error:\n %v %v - Expecting: %v", err, prev, typ)
		}
		return prev
	}
	at := func () tokens.Token { return tokenz[ip]; }
	notEof:=func() bool {if (at().Type == kind.EOF){return false}; return true}

	parse_obj_expr := func () Expr {
		if (at().Type != kind.OpenBrace) {
			return parse_additive_expr()
		}

		next()
		properties := make([]Property, 0)

		for (notEof() && at().Type != kind.ClosedBrace){

			key := expect(
				kind.Identifier,
				"Key Expected For Object",
			).Value

			if (at().Type == kind.Comma){
				next()
				properties = append(properties, Property{Expr{}, kinds.Property{}, key, nil})
				continue
			} else if (at().Type == kind.ClosedBrace) {
				properties = append(properties, Property{Expr{}, kinds.Property{}, key, nil})
				continue
			}

		}

		expect(
			kind.Colon,
			"Missing Colon After Key",
		)
		val:= parse_expr()

		properties = append(properties, Property{Expr{}, kinds.Property{}, key, val})
		if (at().Type != kind.ClosedBrace){
			expect(
				kind.Comma,
				"Comma Or Closing Brace Required After Value",
			)	
		}

		expect(
			kind.ClosedBrace,
			"Object Missing Closeing Brace",
		)	

		return ObjectLiteral{Expr{}, kinds.ObjectLiteral, properties}
	}

	parse_assign_expr := func () Expr {
		left:=parse_obj_expr();

		if (at().Type == kind.Equals){
			next()
			val:=parse_assign_expr()
			return AssignExpr{Expr{}, kinds.AssignExpr{}, left, val}
		}
	}

	parse_expr := func () Expr {
		return parse_assign_expr()
	}


	parse_var_dec := func () Stmnt {

		isConst := (next().Type == kind.Def)
		ident := expect(
			kind.Identifier,
			"Expected Valid Type For Variable Name",
		)

		expect(
			kind.Equals,
			"Expected Equals Symbol To Point Out Var Contents",
		)

		Decl := VarDec{
			kinds.VarDec,
			parse_expr(),
			ident,
			isConst,
		}

		return Decl

	}

	parse_stmnt:=func () Stmnt {
		switch at().Type {
		case kind.Mal:
		case kind.Def:
			return parse_var_dec()

		// default:
		// 	return parse_expr()
		}
	}

	ProdAst := func () Stmnt {
		program := Program{
			kinds.Program,
			make([]Stmnt, 0),
		}

		for (notEof()){
			program.body = append(program.body, parse_stmnt())
		}
	}

}