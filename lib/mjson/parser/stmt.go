package parser

import (
	"fmt"
	"os"
	"soup/lib/mjson/lexer"
	"strconv"
)

func (s *Parser) ParseStmt() Stmt {

	switch s.at().Type {
	case lexer.Int:
		q, _ := strconv.ParseInt(s.next().Value, 10, 64)
		return Int{Type: "Int", Value: q}
	case lexer.Float:
		q, _ := strconv.ParseFloat(s.next().Value, 64)
		return Float{Type: "Float", Value: q}
	case lexer.String:
		return String{Type: "String", Value: s.next().Value}
	case lexer.Null:
		s.next()
		return Null{Type: "Null", Value: nil}
	case lexer.Boolean:
		q := false
		if s.next().Value == "true" {
			q = true
		}
		return Boolean{Type: "Boolean", Value: q}
	case lexer.OpenBracket:
		s.next()
		elements := make([]Stmt, 0)
		for s.at().Type != lexer.ClosedBracket {
			elements = append(elements, s.ParseStmt())
			if s.at().Type == lexer.Comma {
				s.next()
			}
		}
		s.expect(
			lexer.ClosedBracket,
			"Expected Closed Bracket For Array",
		)
		return Array{
			Type: "Array",
			Elements: elements,
		}
	case lexer.OpenBrace:
		s.next()
		properties := make([]Property, 0)
		for s.at().Type != lexer.ClosedBrace {
			key := s.expect(
				lexer.String,
				"Expected String For The Key in a Key:Value Pair",
			).Value
			s.expect(
				lexer.Colon,
				"Expected Colon For Key:Value Pair",
			)
			value := s.ParseStmt()
			properties = append(properties, Property{Type: "Property", Key: key, Value: value})
			if s.at().Type == lexer.Comma {
				s.next()
			}else if s.at().Type == lexer.ClosedBrace {
				s.next()
				break
			}
		}
		return Object{
			Type: "Object",
			Properties: properties,
		}
	default:
		fmt.Println("Invalid Token Found During Parsing", s.at(), "Line:", s.at().Line, "Pos:", s.at().Column)
		os.Exit(1)
		return Null{Type: "Null", Value: nil}

	}

}