package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

// Parser ...
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New ... Creating a new Parser
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{l: lex}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram ... entry point?
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
