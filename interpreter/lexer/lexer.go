package lexer

import (
	"bufio"
	"interpreter/tokens"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	scanner bufio.Reader
	line    int
	column  int
	cc      rune
}

var eofTok = tokens.Token{Type: tokens.EOF}

func (l *Lexer) Source(r io.Reader) {
	l.scanner = *bufio.NewReader(r)
	l.line = 1
	l.column = 0
	l.advance()
}

func (l *Lexer) makeToken(t tokens.Type) tokens.Token {
	return tokens.Token{
		Type:   t,
		Line:   l.line,
		Column: l.column,
	}
}

func (l *Lexer) skipWhitespace() {
	if l.cc == '#' {
		for l.cc != '\n' && l.cc != 0 {
			l.advance()
		}
		if l.cc == '\n' {
			l.advance()
		}
	}
	for unicode.IsSpace(l.cc) && l.cc != 0 {
		l.advance()
	}
}

func (l *Lexer) advance() {
	r, _, err := l.scanner.ReadRune()
	if err != nil {
		l.cc = 0
		return
	}
	l.cc = r
	l.column++
}

func (l *Lexer) number() tokens.Token {
	builder := strings.Builder{}
	for ('0' <= l.cc && l.cc <= '9') || l.cc == '.' || l.cc == '_' {
		if l.cc == 0 {
			break
		}
		builder.WriteRune(l.cc)
		l.advance()
	}
	return tokens.Token{
		Raw:  builder.String(),
		Type: tokens.NUMBER,
	}
}

func (l *Lexer) NextToken() tokens.Token {
	l.skipWhitespace()

	token := tokens.Token{Type: tokens.UNKNOWN}

	switch l.cc {
	case 0:
		token = eofTok
	case '+':
		token = l.makeToken(tokens.PLUS)
	case '-':
		token = l.makeToken(tokens.MINUS)
	case '*':
		token = l.makeToken(tokens.ASTERIKS)
	case '/':
		token = l.makeToken(tokens.SLASH)
	default:
		if '0' <= l.cc && l.cc <= '9' {
			token = l.number()
		}

	}

	l.advance()
	return token
}
