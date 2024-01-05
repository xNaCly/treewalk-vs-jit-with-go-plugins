package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/tokens"
	"log"
	"strconv"
)

type Parser struct {
	l    lexer.Lexer
	cur  tokens.Token
	prev tokens.Token
}

func (p *Parser) Source(l lexer.Lexer) {
	p.l = l
	p.advance()
}

func (p *Parser) Next() ast.Node {
	if !p.atEnd() {
		return p.expression()
	}
	return nil
}

func (p *Parser) expression() ast.Node {
	return p.term()
}

func (p *Parser) term() ast.Node {
	lhs := p.factor()

	for p.match(tokens.MINUS, tokens.PLUS) {
		op := p.prev
		rhs := p.factor()
		lhs = &ast.Binary{
			Token: op,
			Left:  lhs,
			Right: rhs,
		}
	}

	return lhs
}

func (p *Parser) factor() ast.Node {
	lhs := p.unary()

	for p.match(tokens.SLASH, tokens.ASTERIKS) {
		op := p.prev
		rhs := p.unary()
		lhs = &ast.Binary{
			Token: op,
			Left:  lhs,
			Right: rhs,
		}
	}

	return lhs
}

func (p *Parser) unary() ast.Node {
	if p.match(tokens.MINUS) {
		op := p.prev
		right := p.unary()
		return &ast.Unary{Token: op, Right: right}
	}

	return p.primary()
}

func (p *Parser) primary() ast.Node {
	if p.match(tokens.NUMBER) {
		float, err := strconv.ParseFloat(p.prev.Raw, 64)
		if err != nil {
			log.Panic(err)
		}
		return &ast.Number{Token: p.prev, Value: float}
	} else if p.match(tokens.BRACE_LEFT) {
		node := p.expression()
		p.consume(tokens.BRACE_RIGHT, "Expected '('")
		return node
	}

	panic("Expected expression")
}

func (p *Parser) match(tokenTypes ...tokens.Type) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(tokenType tokens.Type, error string) {
	if p.check(tokenType) {
		p.advance()
		return
	}
	log.Panicf("Wanted %q, got %q: %s", tokens.TYPE_MAP[p.cur.Type], tokens.TYPE_MAP[tokenType], error)
}

func (p *Parser) check(tokenType tokens.Type) bool {
	if p.atEnd() {
		return false
	}
	return p.cur.Type == tokenType
}

func (p *Parser) advance() {
	p.prev = p.cur
	p.cur = p.l.Next()
}

func (p *Parser) atEnd() bool {
	return p.cur.Type == tokens.EOF
}
