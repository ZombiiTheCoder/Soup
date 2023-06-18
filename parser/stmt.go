package parser

import (
	"os"
	"path/filepath"
	"soup/ast"
	"soup/tokens"
	"soup/utils"
	"strings"
)

func (s *Parser) ParseStmt() ast.Stmt {

	switch s.Current().Type {
	case tokens.Val, tokens.Var:
		return s.ParseVarDec()
	case tokens.OpenBrace:
		return s.ParseBlock()
	case tokens.If:
		return s.ParseIf()
	case tokens.Function:
		return s.ParseFunction()
	case tokens.While:
		return s.ParseWhile()
	case tokens.Use:
		return s.ParseImport()
	case tokens.Return:
		return s.ParseReturn()
	default:
		return s.ParseExpr()
	}

}

func (s *Parser) ParseVarDec() ast.Stmt {

	// type id = expr;
	NotMut := s.Eat().Type == tokens.Val
	id := s.Expect(
		tokens.Identifier,
		"Identifier Required For Declaring Variable",
	).Value

	if s.Current().Type == tokens.Semicolon {
		s.Eat()
		if NotMut {
			utils.Error("Value Must Be Assigned To Constant")
		}
		
		return ast.VarDec{
			Type: "VarDec",
			Name: id,
			Value: nil,
			NotMut: NotMut,
		}
	}
	s.Expect(
		tokens.Equals,
		"Equal Sign Required For Assigning Value To Variable",
	)

	dec := ast.VarDec{
		Type: "VarDec",
		Name: id,
		Value: s.ParseExpr(),
		NotMut: NotMut,
	}
	s.EatSemi()
	return dec


}

func (s *Parser) ParseBlock() ast.BlockStmt {

	stmts := make([]ast.Stmt, 0)
	s.Expect(
		tokens.OpenBrace,
		"Expected Open Brace For Start Of Block",
	)

	for s.NotEof() && s.Current().Type != tokens.ClosedBrace {
		stmts = append(stmts, s.ParseStmt())
	}

	s.Expect(
		tokens.ClosedBrace,
		"Expected Closed Brace For Start Of Block",
	)
	s.EatSemi()
	return ast.BlockStmt{
		Type: "BlockStmt",
		Body: stmts,
	}

}

func (s *Parser) ParseIf() ast.Stmt {

	s.Eat()
	s.Expect(
		tokens.OpenParen,
		"Expected Open Paren For If Statement",
	)

	condition := s.ParseStmt()

	s.Expect(
		tokens.ClosedParen,
		"Expected Closing Paren For If Statement",
	)

	consequent := s.ParseBlock().Body

	var alternate any
	if s.Current().Type == tokens.Else {
		s.Eat()
		if s.Current().Type == tokens.If {
			alternate = s.ParseIf()
		}else{
			alternate = s.ParseBlock().Body
		}
	}

	return ast.IfStmt{
		Type: "IfStmt",
		Test: condition,
		Consquent: consequent,
		Alternate: alternate,
	}

}

func (s *Parser) ParseWhile() ast.Stmt {

	s.Eat()
	s.Expect(
		tokens.OpenParen,
		"Expected Open Paren For While Statement",
	)

	condition := s.ParseStmt()

	s.Expect(
		tokens.ClosedParen,
		"Expected Closing Paren For While Statement",
	)

	consequent := s.ParseBlock().Body

	return ast.WhileStmt{
		Type: "WhileStmt",
		Test: condition,
		Consquent: consequent,
	}

}

func (s *Parser) ParseImport() ast.Stmt {

	s.Eat()

	ex, _ := os.Executable()
	realEx, _ := filepath.EvalSymlinks(ex)
	ExeLocation := filepath.Dir(realEx)
	StdPath := filepath.Join(ExeLocation, "/pkg/")
	q, _ := filepath.Abs(s.FileName)
	FilePath := filepath.Dir(q)

	F := s.Expect(
		tokens.String,
		"String Expected For Import Statement",
	).Value

	file := ""
	if !strings.Contains(F, ".soup") {
		file = filepath.Join(F, "/init.soup")
	} else {
		file = F
	}

	f := ""
	rel := !strings.Contains(F, "@") || !strings.Contains(F, "pkg:")
	if strings.Contains(F, "@") || strings.Contains(F, "pkg:") {
		f = filepath.Join(StdPath, strings.ReplaceAll(file, strings.ReplaceAll(file, "pkg:", ""), ""))
	} else {
		f = filepath.Join(FilePath, file)
	}

	if strings.Contains(F, "@") && strings.Contains(F, "pkg:") {
		utils.Error("Import contains `@` symbol and `pkg:` path identifier")
	}

	return ast.ImpStmt{
		Type: "ImpStmt",
		File: strings.ReplaceAll(f, "\\", "/"),
		Rel: rel,
	}

}

func (s *Parser) ParseFunction() ast.Stmt {

	s.Eat()
	name := s.Expect(
		tokens.Identifier,
		`Name For Function Expected following fn or func`,
	).Value

	args := s.ParseArgs()
	params := make([]string, 0)

	for _, v := range args {
		if v.GetType() != "Identifier" {
			utils.Error("\nParemeter For Function Must Be Identifier")
		}

		params = append(params, v.(ast.Identifier).Symb)
	}

	body := s.ParseBlock().Body
	s.EatSemi()

	return ast.FuncDec{
		Type: "FuncDec",
		Name: name,
		Params: params,
		Body: body,
	}

}

func (s *Parser) ParseReturn() ast.Stmt {

	s.Eat()
	val := s.ParseExpr()
	s.EatSemi()

	return ast.ReturnStmt{
		Type: "ReturnStmt",
		Value: val,
	}

}