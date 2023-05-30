package parser

import (
	"Soup/src/ast"
	"Soup/src/lex2"
	"strings"

	// "Soup/src/lexer"
	"Soup/src/lex2/token"
	"Soup/src/lex2/token/kind"
	"Soup/src/utils/fmt"
	f "fmt"
	"os"
)

type Parser struct {
	Ip     int
	Tokens []token.Token
	Src    string
}

func (s *Parser) At() token.Token {

	return s.Tokens[s.Ip]

}

func (s *Parser) Eat() token.Token {

	s.Ip++
	return s.Tokens[s.Ip-1]

}

func (s *Parser) Expect(ExType kind.TokenType, err string) token.Token {

	s.Ip++
	last := s.Tokens[s.Ip-1]
	if last.Type != ExType {
		f.Printf("\nParser Error:\n %v - %v Expecting %v\n", err, last, ExType)
		os.Exit(1)
	}

	return last

}

func (s *Parser) Not_Eof() bool {

	return s.Tokens[s.Ip].Type != kind.EOF

}

func (s *Parser) Pass_Semi_Colon() {
	if s.At().Type == kind.Semicolon {
		s.Eat()
	}
}

func (s *Parser) ProdAst() ast.Stmt {
	s.Tokens = lex2.BuildLexer(s.Src)
	prg := ast.Create_Program(make([]ast.Stmt, 0)).(ast.Program)

	for s.Not_Eof() {
		prg.Body = append(prg.Body, s.parse_stmt())
	}

	return prg

}

func (s *Parser) parse_var_dec() ast.Stmt {

	isConst := (s.Eat().Type == kind.Val)
	ident := s.Expect(
		kind.Identifier,
		"Expected Identifier name following def or var",
	).Value

	if s.At().Type == kind.Semicolon {
		s.Eat()
		if isConst {
			fmt.Prints.Error("Must Assign Value to Def")
		}

		return ast.Create_VarDec(isConst, ident, nil)
	}

	s.Expect(
		kind.Equals,
		"Equals Token Expected",
	)

	dec := ast.Create_VarDec(isConst, ident, s.parse_expr())

	s.Pass_Semi_Colon()

	return dec

}

func (s *Parser) parse_fn_dec() ast.Stmt {
	s.Eat()
	name := s.Expect(
		kind.Identifier,
		`Name For Function Expected following fn or func`,
	).Value

	args := s.parse_args()
	params := make([]string, 0)
	for _, v := range args {
		if v.GetType() != "Identifier" {
			f.Printf("\nParemeter For Function Must Be Identifier %v\n", v.GetValue())
		}

		params = append(params, v.(ast.Identifier).Symb)
	}

	s.Expect(
		kind.OpenBrace,
		"Expected Function Body Following Declaration",
	)

	body := make([]ast.Stmt, 0)

	for s.At().Type != kind.EOF && s.At().Type != kind.ClosedBrace {
		body = append(body, s.parse_stmt())
	}

	s.Expect(
		kind.ClosedBrace,
		"Closing Brace Expected For Function Body",
	)

	fn := ast.Create_Function(name, params, body)
	return fn
}

func (s *Parser) parse_ret_stmt() ast.Stmt {
	s.Eat()
	Val := s.parse_stmt()

	return ast.Create_Ret_Stmt(Val)
}

func (s *Parser) parse_Import_stmt() ast.Stmt {
	s.Eat()
	ffl := s.Expect(
		kind.String,
		"String Expected For Import Value",
	).Value

	return ast.Create_Import_Stmt(ffl, !strings.Contains(ffl, "@"))
}

func (s *Parser) parse_block_stmt() ast.BlockStmt {

	Stmts := make([]ast.Stmt, 0)
	s.Expect(kind.OpenBrace, "Expected Opening Brace For The Block")

	for s.Not_Eof() && s.At().Type != kind.ClosedBrace {
		Stmts = append(Stmts, s.parse_stmt())
	}

	s.Expect(kind.ClosedBrace, "Expected Closing Brace For The Block")

	return ast.Create_Block_Stmt(
		Stmts,
	)
}

func (s *Parser) parse_if_stmt() ast.Stmt {
	s.Eat()
	s.Expect(kind.OpenParen, "Expected Open Paren for if Statement")
	condition := s.parse_stmt()
	s.Expect(kind.ClosedParen, "Expected Closed Paren for if Statement")

	consequent := s.parse_block_stmt()

	var alternate any
	if s.At().Type == kind.Else {
		s.Eat()
		if s.At().Type == kind.If {
			alternate = s.parse_if_stmt()
		} else {
			alternate = s.parse_block_stmt()
		}
	}

	return ast.Create_If_Stmt(
		condition,
		consequent,
		alternate,
	)
}

func (s *Parser) parse_while_stmt() ast.Stmt {
	s.Eat()
	s.Expect(kind.OpenParen, "Expected Open Paren for while Statement")
	condition := s.parse_stmt()
	s.Expect(kind.ClosedParen, "Expected Closed Paren for while Statement")

	consequent := s.parse_block_stmt()

	return ast.Create_While_Stmt(
		condition,
		consequent,
	)
}

func (s *Parser) parse_stmt() ast.Stmt {

	switch s.At().Type {

	case kind.Val, kind.Var:
		return s.parse_var_dec()
	case kind.If:
		return s.parse_if_stmt()
	case kind.While:
		return s.parse_while_stmt()
	case kind.OpenBrace:
		return s.parse_block_stmt()
	case kind.Fn:
		return s.parse_fn_dec()
	case kind.Ret:
		return s.parse_ret_stmt()
	case kind.Use:
		return s.parse_Import_stmt()
	default:
		return s.parse_expr()

	}

}

func (s *Parser) parse_expr() ast.Expr {

	return s.parse_assign_expr()

}

func (s *Parser) parse_assign_expr() ast.Expr {

	left := s.parse_relational_expr()

	if s.At().Type == kind.Equals ||
		s.At().Type == kind.PlusEquals ||
		s.At().Type == kind.MinusEquals ||
		s.At().Type == kind.DivideEquals ||
		s.At().Type == kind.PercentEquals ||
		s.At().Type == kind.MultiplyEquals {

		op := s.Eat().Value
		val := s.parse_assign_expr()

		return ast.Create_AssignExpr(left, val, op)
	}

	return left

}

func (s *Parser) parse_relational_expr() ast.Expr {

	left := s.parse_object_expr()

	for s.At().Value == "!=" || s.At().Value == "==" || s.At().Value == "<" || s.At().Value == "<=" || s.At().Value == ">" || s.At().Value == ">=" {

		op := s.Eat().Value
		right := s.parse_object_expr()

		left = ast.Create_Relational_Expr(left, right, op)

	}

	return left
}

func (s *Parser) parse_object_expr() ast.Expr {
	if s.At().Type != kind.OpenBrace {
		return s.parse_additive_expr()
	}

	s.Eat() // advance past open brace.
	properties := make([]ast.Property, 0)

	for s.Not_Eof() && s.At().Type != kind.ClosedBrace {
		key := s.Expect(kind.Identifier, "Object literal key exprected").Value

		if s.At().Type == kind.Comma {
			s.Eat()
			properties = append(properties, ast.Create_Property(key, nil).(ast.Property))
			continue
		} else if s.At().Type == kind.ClosedBrace {
			properties = append(properties, ast.Create_Property(key, nil).(ast.Property))
			continue
		}

		s.Expect(
			kind.Colon,
			"Missing colon following identifier in ObjectExpr",
		)
		value := s.parse_expr()
		properties = append(properties, ast.Create_Property(key, value).(ast.Property))

		if s.At().Type != kind.ClosedBrace {
			s.Expect(
				kind.Comma,
				"Expected comma or closing bracket following property",
			)
		}
	}

	s.Expect(kind.ClosedBrace, "Object literal missing closing brace.")
	return ast.Create_ObjectLiteral(properties)
}

func (s *Parser) parse_additive_expr() ast.Expr {

	left := s.parse_multiplicative_expr()
	for s.At().Value == "+" || s.At().Value == "-" {
		op := s.Eat().Value
		right := s.parse_multiplicative_expr()

		left = ast.Create_BinaryExpr(left, right, op)

	}

	return left

}

func (s *Parser) parse_multiplicative_expr() ast.Expr {

	left := s.parse_unary_expr()

	for s.At().Value == "/" || s.At().Value == "*" || s.At().Value == "%" {

		op := s.Eat().Value
		right := s.parse_unary_expr()

		left = ast.Create_BinaryExpr(left, right, op)

	}

	return left

}

func (s *Parser) parse_unary_expr() ast.Expr {

	var left any
	if s.At().Value == "--" || s.At().Value == "++" || s.At().Value == "+" || s.At().Value == "-" || s.At().Value == "~" || s.At().Value == "!" {
		for s.At().Value == "--" || s.At().Value == "++" || s.At().Value == "+" || s.At().Value == "-" || s.At().Value == "~" || s.At().Value == "!" {
			op := s.Eat().Value
			left = s.parse_call_member_expr()
			left = ast.Create_UnaryExpr(left.(ast.Expr), op, true)
		}
	} else {
		left = s.parse_call_member_expr()
		if s.At().Value == "--" || s.At().Value == "++" {
			for s.At().Value == "--" || s.At().Value == "++" {
				op := s.Eat().Value
				left = ast.Create_UnaryExpr(left.(ast.Expr), op, false)
			}

		}
	}
	return left.(ast.Expr)

}

func (s *Parser) parse_array_expr() ast.Expr {

	s.Eat()
	var elements []ast.Expr = make([]ast.Expr, 0)

	for s.At().Type != kind.ClosedBracket {
		elements = append(elements, s.parse_expr())
		if s.At().Type == kind.Comma {
			s.Eat()
		}

	}

	s.Expect(
		kind.ClosedBracket,
		"Missing Closing Bracket For Array",
	)

	return ast.Create_Array_Expr(
		elements,
	)

}

func (s *Parser) parse_call_member_expr() ast.Expr {

	member := s.parse_member_expr()

	if s.At().Type == kind.OpenParen {
		return s.parse_call_expr(member)
	}

	if s.At().Type == kind.OpenBracket {
		return s.parse_array_expr()
	}

	return member

}

func (s *Parser) parse_call_expr(caller ast.Expr) ast.Expr {

	call_expr := ast.Create_CallExpr(
		caller,
		s.parse_args(),
	)

	if s.At().Type == kind.OpenParen {
		call_expr = s.parse_call_expr(call_expr)
	}

	return call_expr
}

func (s *Parser) parse_args() []ast.Expr {
	s.Expect(
		kind.OpenParen,
		"Expected Open Paren For Args List",
	)

	var args []ast.Expr
	if s.At().Type == kind.ClosedParen {
		args = make([]ast.Expr, 0)
	} else {
		args = s.parse_args_list()
	}

	s.Expect(
		kind.ClosedParen,
		"Missing Close Paren For Args List",
	)
	return args
}

func (s *Parser) parse_args_list() []ast.Expr {
	args := make([]ast.Expr, 0)
	args = append(args, s.parse_assign_expr())

	for s.At().Type == kind.Comma {
		s.Eat()
		args = append(args, s.parse_assign_expr())
	}

	return args
}

func (s *Parser) parse_member_expr() ast.Expr {

	obj := s.parse_primary_expr()

	for s.At().Type == kind.Period {
		op := s.Eat()
		var property ast.Expr
		var computed bool

		if op.Type == kind.Period {
			computed = false
			// get identifier
			property = s.parse_primary_expr()
			if property.GetType() != "Identifier" {
				fmt.Prints.Error("Cannot use dot operator without right hand side being a identifier")
			}
		} else if s.At().Type == kind.OpenBracket { // this allows obj[computedValue]
			s.Eat()
			computed = true
			property = s.parse_expr()
			s.Expect(
				kind.ClosedBracket,
				"Missing closing bracket in computed value.",
			)
		}
		obj = ast.Create_MemberExpr(obj, property, computed)
	}

	return obj

}

func (s *Parser) parse_primary_expr() ast.Expr {

	tk := s.At().Type

	switch tk {
	case kind.Identifier:
		return ast.Create_Identifier(s.Eat().Value)

	case kind.Numeral:
		return ast.Create_NumericLiteral(s.Eat().Value)

	case kind.Float:
		return ast.Create_FloatLiteral(s.Eat().Value)

	case kind.String:
		return ast.Create_StringLiteral(s.Eat().Value)

	case kind.Exclamation, kind.Tilde, kind.Plus, kind.Minus, kind.DPlus, kind.DMinus:
		return s.parse_unary_expr()

	case kind.OpenParen:
		s.Eat()
		value := s.parse_expr()
		s.Expect(
			kind.ClosedParen,
			"Unexpected token found inside parenthesised expression. Expected closing parenthesis.",
		)

		return value

	case kind.OpenBracket:
		return s.parse_array_expr()

	default:
		f.Printf("\nUnexpected token found during parsing! %v", s.At())
		os.Exit(1)
	}

	return ast.Create_StringLiteral("")

}

func BuildParser(text string) ast.Stmt {
	prse := Parser{Src: text}
	return prse.ProdAst()
}
