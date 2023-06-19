package parser

import (
	"soup/ast"
	"soup/tokens"
	"soup/utils"
	"strconv"
)

func (s *Parser) ParseExpr() ast.Expr {

	value := s.ParseAssign()
	s.EatSemi()
	return value

}

func (s *Parser) ParseAssign() ast.Expr {

	left := s.ParseTernary()

	if s.Current().Type == tokens.Equals ||
	s.Current().Type == tokens.PlusEquals ||
	s.Current().Type == tokens.DashEquals ||
	s.Current().Type == tokens.StarEquals ||
	s.Current().Type == tokens.ModuloEquals ||
	s.Current().Type == tokens.SlashEquals {

		op := s.Eat().Value
		value := s.ParseAssign()

		left = ast.AssignExpr{
			Type: "AssignExpr",
			Assigner: left,
			Op: op,
			Val: value,
		}

	}

	return left

}

func (s *Parser) ParseTernary() ast.Expr {

	left := s.ParseLogicalOr()

	for s.Current().Type == tokens.QuestionMark {
		s.Eat()
		consequent := s.ParseLogicalOr()
		var alternate interface{}
		if (s.Current().Type == tokens.Colon) {
			s.Eat()
			alternate = s.ParseTernary()
		}else{
			utils.Error(
				"Else for Ternary Expression Missing\nExpected Token %v Found Token %v at line %v column %v\n Token Location %v:%v:%v",
				":",
				s.Current().Type,
				s.Current().Line,
				s.Current().Column,
				s.Current().FileName,
				s.Current().Line,
				s.Current().Column,
			)
		}
		left = ast.TernaryExpr{
			Type: "TernaryExpr",
			Condition: left,
			Consquent: consequent,
			Alternate:  alternate,
		}

	}

	return left

}

func (s *Parser) ParseLogicalOr() ast.Expr {
	left := s.ParseLogicalAnd()

	for s.Current().Type == tokens.Or {
		op := s.Eat().Value
		right := s.ParseLogicalAnd()

		left = ast.LogicalExpr{
			Type: "LogicalExpr",
			Op: op,
			Left: left,
			Right: right,
		}

	}

	return left
}

func (s *Parser) ParseLogicalAnd() ast.Expr {
	
	left := s.ParseBitwiseOr()
	for s.Current().Type == tokens.And {
		op := s.Eat().Value
		right := s.ParseBitwiseOr()

		left =  ast.LogicalExpr{
			Type: "LogicalExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left
}

func (s *Parser) ParseBitwiseOr() ast.Expr {
	left := s.ParseBitwiseAnd()
	for s.Current().Type == tokens.Pipe {
		op := s.Eat().Value
		right := s.ParseBitwiseAnd()

		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left
}

func (s *Parser) ParseBitwiseAnd() ast.Expr {
	left := s.ParseEquality()
	for s.Current().Type == tokens.Ampersan {
		op := s.Eat().Value
		right := s.ParseEquality()

		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left
}

func (s *Parser) ParseEquality() ast.Expr {
	left := s.ParseRelational()
	for s.Current().Type == tokens.NEquals||
	s.Current().Type == tokens.DEquals {
		op := s.Eat().Value
		right := s.ParseRelational()

		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left
}

func (s *Parser) ParseRelational() ast.Expr {
	left := s.ParseObject()
	for s.Current().Type == tokens.LessThan||
	s.Current().Type == tokens.LTEquals ||
	s.Current().Type == tokens.GreaterThan||
	s.Current().Type == tokens.GTEquals {
		op := s.Eat().Value
		right := s.ParseObject()

		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left
}

func (s *Parser) ParseObject() ast.Expr {

	if s.Current().Type != tokens.OpenBrace {
		return s.ParseAdditive()
	}

	s.Eat()
	properties := make([]ast.Property, 0)
	
	for s.NotEof() && s.Current().Type != tokens.ClosedBrace {

		key := s.Expect(tokens.Identifier, "Object key expected").Value
		if s.Current().Type == tokens.Comma {
			s.Eat()
			properties = append(properties, ast.Property{
				Type: "Property",
				Key: key,
				Val: nil,
			})
			continue
			} else if s.Current().Type == tokens.ClosedBrace {
				properties = append(properties, ast.Property{
					Type: "Property",
					Key: key,
					Val: nil,
				})
				continue
			}
			
		s.Expect(
			tokens.Colon,
			"Missing colon following identifier in Object",
		)
		val := s.ParseExpr()
		


		properties = append(properties, ast.Property{
			Type: "Property",
			Key: key,
			Val: val,
		})

		if s.Current().Type == tokens.Comma {
			s.Eat()
		}

	}

	if s.Current().Type != tokens.ClosedBrace {
		s.Expect(
			tokens.Comma,
			"Expected comma or closing bracket following property",
		)
	}

	s.Expect(tokens.ClosedBrace, "Object missing closing brace.")
	return ast.ObjectLiteral{
		Type: "ObjectLiteral",
		Properties: properties,
	}

}

func (s *Parser) ParseAdditive() ast.Expr {

	left := s.ParseMultiplicative()
	for s.Current().Type == tokens.Plus ||
	s.Current().Type == tokens.Minus {
		op := s.Eat().Value
		right := s.ParseMultiplicative()
		
		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}

	return left

}

func (s *Parser) ParseMultiplicative() ast.Expr {

	left := s.ParseUnaryPrefix()
	for s.Current().Type == tokens.Star ||
	s.Current().Type == tokens.Slash ||
	s.Current().Type == tokens.Percent {
		op := s.Eat().Value
		right := s.ParseUnaryPrefix()
		
		left = ast.BinaryExpr{
			Type: "BinaryExpr",
			Op: op,
			Left: left,
			Right: right,
		}
	}
	return left

}

func (s *Parser) ParseUnaryPrefix() ast.Expr {

	var left any
	if s.Current().Type == tokens.Plus ||
	s.Current().Type == tokens.Minus ||
	s.Current().Type == tokens.Bang ||
	s.Current().Type == tokens.Tilde ||
	s.Current().Type == tokens.DPlus ||
	s.Current().Type == tokens.DMinus {
		op := s.Eat().Value
		if s.Current().Type == tokens.DPlus || s.Current().Type == tokens.DMinus && s.Current().Type != tokens.Identifier{
			q := s.Eat()
			utils.Error(
				"Expected %v Found Token %v at line %v column %v\n Token Location %v:%v:%v",
				tokens.Identifier,
				q.Type,
				q.Line,
				q.Column,
				q.FileName,
				q.Line,
				q.Column,
			)
		}
		left = s.ParseUnaryPostfix()
		return ast.UnaryExpr{
			Type: "UnaryExpr",
			Op: op,
			Argument: left.(ast.Expr),
			Prefix: true}
	}

	if (left != nil){
		return left.(ast.Expr)
	}else{
		return s.ParseUnaryPostfix()
	}
}

func (s *Parser) ParseUnaryPostfix() ast.Expr {
	tpe := s.Current()
	var left = s.ParseCallMember()
	if s.Current().Type == tokens.DPlus ||
	s.Current().Type == tokens.DMinus {
		op := s.Eat().Value
		if tpe.Type != tokens.Identifier{
			utils.Error(
				"Expected %v Found Token %v at line %v column %v\n Token Location %v:%v:%v",
				tokens.Identifier,
				tpe.Type,
				tpe.Line,
				tpe.Column,
				tpe.FileName,
				tpe.Line,
				tpe.Column,
			)
		}
		left = ast.UnaryExpr{
			Type: "UnaryExpr",
			Op: op,
			Argument: left,
			Prefix: false,
		}
	}
	return left
}

func (s *Parser) ParseCallMember() ast.Expr {

	member := s.ParseMember()

	if s.Current().Type == tokens.OpenParen {
		return s.ParseCall(member)
	}

	return member

}

func (s *Parser) ParseCall(caller ast.Expr) ast.Expr {

	callexpr := ast.CallExpr{
		Type: "CallExpr",
		Caller: caller,
		Args: s.ParseArgs(),
	}

	if s.Current().Type == tokens.OpenParen {
		callexpr = s.ParseCall(callexpr).(ast.CallExpr)
	}

	return callexpr

}

func (s *Parser) ParseArgs() []ast.Expr {

	s.Expect(
		tokens.OpenParen,
		"Expected Open Paren For List",
	)

	var args []ast.Expr
	if s.Current().Type == tokens.ClosedParen {
		args = make([]ast.Expr, 0)
	} else {
		args = s.ParseArgsList()
	}

	s.Expect(
		tokens.ClosedParen,
		"Expected Closing Paren For List",
	)

	return args

}

func (s *Parser) ParseArgsList() []ast.Expr {

	args := make([]ast.Expr, 0)
	args = append(args, s.ParseAssign())

	for s.Current().Type == tokens.Comma {
		s.Eat()
		args = append(args, s.ParseAssign())
	}

	return args

}

func (s *Parser) ParseMember() ast.Expr {

	obj := s.ParsePrimary()

	for s.Current().Type == tokens.Dot || s.Current().Type == tokens.OpenBracket {
		op := s.Eat()
		var property ast.Expr
		var computed bool

		if op.Type == tokens.Dot {
			computed = false

			property = s.ParsePrimary()
			if property.GetType() != "Identifier" {
				utils.Error("Cannot use dot without right side being identifier")
			}

		}else if op.Type == tokens.OpenBracket {

			// s.Eat()
			computed = true
			property = s.ParseExpr()
			s.Expect(
				tokens.ClosedBracket,
				"Missing Closing Bracket in Computed Value",
			)

			
		}
		obj = ast.MemberExpr{
			Type: "MemberExpr",
			Obj: obj,
			Property: property,
			Computed: computed,
		}

	}

	return obj

}

func (s *Parser) ParseArray() ast.Expr {

	s.Eat()
	var elements = make([]ast.Expr, 0)

	for s.Current().Type != tokens.ClosedBracket {
		elements = append(elements, s.ParseExpr())
		if s.Current().Type == tokens.Comma {
			s.Eat()
		}
	}

	s.Expect(
		tokens.ClosedBracket,
		"Missing Closing Bracket For Array",
	)

	s.EatSemi()

	return ast.ArrayExpr{
		Type: "ArrayExpr",
		Elements: elements,
	}

}

func (s *Parser) ParsePrimary() ast.Expr {

	switch s.Current().Type{

	case tokens.Identifier:
		return ast.Identifier{
		Type: "Identifier",
		Symb: s.Eat().Value,
	}

	case tokens.Number:
		q, _ := strconv.ParseInt(s.Eat().Value, 10, 64)
		return ast.IntegerLiteral{
		Type: "IntegerLiteral",
		Valu: q,
	}

	case tokens.Float:
		q, _ := strconv.ParseFloat(s.Eat().Value, 64)
		return ast.FloatLiteral{
		Type: "FloatLiteral",
		Valu: q,
	}

	case tokens.String:
		return ast.StringLiteral{
		Type: "StringLiteral",
		Valu: s.Eat().Value,
	}
	
	case tokens.OpenParen:
		s.Eat()
		val := s.ParseExpr()
		s.Expect(
			tokens.ClosedParen,
			"Closing Paren Is Missing For Parenthesised Expr",
		)
		s.EatSemi()
		return val

	case tokens.OpenBracket:
		return s.ParseArray()

	default:
		utils.Error("Invalid Token Found During Parsing %v\n %v:%v:%v", s.Current(),
		s.Tokens[s.I].FileName,
		s.Tokens[s.I].Line,
		s.Tokens[s.I].Column,
	)
		return ast.Identifier{}

	}

}