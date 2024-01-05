package lexer

import (
	"interpreter/tokens"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerWhitespace(t *testing.T) {
	tests := []string{
		" ",
		"\n",
		" \n",
		"\t",
		"# comment\n ",
		`
        #comment 1
        #comment 2
        # comment 3
        `,
	}
	l := Lexer{}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			l.Source(strings.NewReader(test))
			res := l.Next()
			if res.Type != tokens.EOF {
				t.Errorf("Token type incorrect, wanted eof, got: %#+v", res)
			}
		})
	}
}

func TestLexerSymbols(t *testing.T) {
	input := "+-/*()"
	expected := []tokens.Type{tokens.PLUS, tokens.MINUS, tokens.SLASH, tokens.ASTERIKS, tokens.BRACE_LEFT, tokens.BRACE_RIGHT, tokens.EOF}
	l := Lexer{}
	l.Source(strings.NewReader(input))

	for _, exp := range expected {
		assert.Equal(t, exp, l.Next().Type)
	}
}

func TestLexerDigit(t *testing.T) {
	tests := []string{
		"125_120.0",
		"25.0",
		"12_000_000",
	}
	l := Lexer{}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			l.Source(strings.NewReader(test))
			res := l.Next()
			if res.Type != tokens.NUMBER {
				t.Errorf("Token type incorrect, wanted number, got: %#+v", res)
			}
			if res.Raw != test {
				t.Errorf("Token value incorrect, wanted %q, got: %q", test, res.Raw)
			}
		})
	}
}

func TestLexerComplete(t *testing.T) {
	tests := []struct {
		in  string
		exp []tokens.Type
	}{
		{"1+1", []tokens.Type{tokens.NUMBER, tokens.PLUS, tokens.NUMBER, tokens.EOF}},
	}
	l := Lexer{}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			l.Source(strings.NewReader(test.in))
			for _, e := range test.exp {
				assert.Equal(t, e, l.Next().Type)
			}
		})
	}
}
