package parser

import (
	"github.com/lerity-yao/cztctl/goTemplate/rabbitmqGen/scanner"
	"github.com/lerity-yao/cztctl/goTemplate/rabbitmqGen/token"
	"log"
	"path/filepath"
)

const (
	AT_TYPE    = "type"
	AT_SERVER  = "@server"
	AT_SERVICE = "service"
)

type Parser struct {
	s      *scanner.Scanner
	errors []error

	curTok     token.Token
	peekTok    token.Token
	node       []token.Token
	curTokStmt string
}

// New creates a new parser.
func New(filename string, src interface{}) *Parser {
	abs, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalln(err)
	}
	//
	p := &Parser{
		s: scanner.MustNewScanner(abs, src),
	}

	return p
}

// Parse parses the api file.
func (p *Parser) Parse() error {
	if !p.init() {
		return nil
	}

	// 收集所有token
	for p.curTokenIsNotEof() {
		p.loadTokStmt()
		if !p.nextToken() {
			return nil
		}
	}

	p.parseStmt()

	return nil
}

func (p *Parser) init() bool {
	if !p.nextToken() {
		return false
	}
	return p.nextToken()
}

func (p *Parser) nextToken() bool {
	var err error
	p.curTok = p.peekTok
	if p.curTok.Valid() {
		if p.curTokenIs(token.EOF) {
			return false
		}
		p.peekTok, err = p.s.NextToken()
		if err != nil {
			p.errors = append(p.errors, err)
			return false
		}

		p.node = append(p.node, p.curTok)
		return true
	}

	p.peekTok, err = p.s.NextToken()
	if err != nil {
		p.errors = append(p.errors, err)
		return false
	}
	p.node = append(p.node, p.curTok)
	return true
}

func (p *Parser) curTokenIsNotEof() bool {
	return p.curTokenIsNot(token.EOF)
}

func (p *Parser) loadTokStmt() {
	switch p.curTok.Text {
	case AT_TYPE:
		p.curTokStmt = AT_TYPE
	case AT_SERVICE:
		p.curTokStmt = AT_SERVICE
	case AT_SERVER:
		p.curTokStmt = AT_SERVER
	default:
	}
}

func (p *Parser) curTokenIsNot(expected token.Type) bool {
	return p.curTok.Type != expected
}

func (p *Parser) curTokenIs(expected ...interface{}) bool {
	for _, v := range expected {
		switch val := v.(type) {
		case token.Type:
			if p.curTok.Type == val {
				return true
			}
		case string:
			if p.curTok.Text == val {
				return true
			}
		}
	}

	return false
}

// parseStmt 解析语句
func (p *Parser) parseStmt() {
	//for idx, item := range p.node {
	//	fmt.Println(idx, item)
	//	switch item.Text {
	//	case AT_TYPE:
	//		p.parseType(item)
	//	case AT_SERVER:
	//
	//	case AT_SERVICE:
	//
	//	default:
	//		fmt.Println("default")
	//		//p.errors = append(p.errors, fmt.Errorf("%s unexpected token '%s'", p.curTok.Position.String(), p.peekTok.Text))
	//	}
	//	//
	//}
}

func (p *Parser) curTokenTextIs(expected string) bool {
	if p.curTok.Text == expected {
		return true
	}
	return false
}
