package parser

import (
	"fmt"
	"os"
	"soup/lexer"
	. "soup/parser/ast"
)

func (s *Parser) parseStmt() Stmt {
	switch s.at().Type {
	case lexer.Const, lexer.Var:
		return s.parseVarDeclaration()
	case lexer.OpenBrace:
		return s.parseBlock()
	case lexer.If:
		return s.parseIf()
	case lexer.Function:
		return s.parseFunction()
	case lexer.While:
		return s.parseWhile()
	case lexer.Return:
		return s.parseReturn()
	default:
		return s.parseExpr()
	}
}

func (s *Parser) parseVarDeclaration() Stmt {

	Const := s.next().Type == lexer.Const
	name := s.expect(lexer.Identifier).Value
	Type := s.expect(lexer.Type).Value
	if s.at().Type == lexer.SemiColon {
		s.next()
		if Const {
			fmt.Println("Value", name, "must be assigned to the Const keyword")
			os.Exit(1)
		}

		return VarDeclaration{
			NodeType: "VarDec",
			Type: Type,
			Name: name,
			Value: nil,
		}
	}
	s.expect(lexer.Equals)

	Value := s.parseExpr()
	s.eatSemi()
	return VarDeclaration{
		NodeType: "VarDec",
		Type: Type,
		Name: name,
		Value: Value,
		Const: Const,
	}

}

func (s *Parser) parseBlock() Stmt {
	
	stmts := make([]Stmt, 0)
	s.expect(lexer.OpenBrace)
	for s.notEof() && s.at().Type != lexer.ClosedBrace {
		stmts = append(stmts, s.parseStmt())
	}

	s.expect(lexer.ClosedBrace)
	s.eatSemi()
	return BlockStmt{
		NodeType: "BlockStmt",
		Type: "null",
		Body: stmts,
	}

}

func (s *Parser) parseIf() Stmt {
	s.next()
	s.expect(lexer.OpenParen)
	condition := s.parseStmt()
	s.expect(lexer.ClosedParen)
	consequent := s.parseBlock().(BlockStmt).Body
	var alternate any
	if s.at().Type == lexer.Else {
		s.next()
		if s.at().Type == lexer.If {
			alternate = s.parseIf()
		}else {
			alternate = s.parseBlock().(BlockStmt).Body
		}
	}

	return IfStmt{
		NodeType: "IfStmt",
		Type: "null",
		Condition: condition,
		Consequent: consequent,
		Alternate: alternate,
	}
}

func (s *Parser) parseFunction() Stmt {
	s.next()
	name := s.expect(lexer.Identifier).Value
	args := s.ParseTypedArgs()
	params := make(map[string]string)

	for k, v := range args {
		if k.GetType() != "Type" && v.GetType() != "Identifier" {
			fmt.Println("Parameter For Function",name, "Must Be Identifier")
		}

		params[k.(Identifier).Value] = v.(Identifier).Value
	}
	body := s.parseBlock().Body
	s.eatSemi()

	return FuncDec{
		Type: "FuncDec",
		Name: name,
		Params: params,
		Body: body,
	}
}

func (s *Parser) parseWhile() Stmt {
	
}

func (s *Parser) parseReturn() Stmt {
	
}