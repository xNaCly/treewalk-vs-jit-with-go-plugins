package lexer

import (
	"interpreter/tokens"
	"strings"
	"testing"
)

func TestLexerWhitespace(t *testing.T) {
	tests := []string{
		" ",
		"\n",
		" \n",
		"\t",
		"# comment\n ",
	}
	l := Lexer{}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			l.Source(strings.NewReader(test))
			res := l.NextToken()
			if res.Type != tokens.EOF {
				t.Errorf("Token type incorrect, wanted eof, got: %#+v", res)
			}
		})
	}
}

func TestLexerSymbols(t *testing.T) {
	input := "+-/*"
	expected := []tokens.Type{tokens.PLUS, tokens.MINUS, tokens.SLASH, tokens.ASTERIKS, tokens.EOF}
	l := Lexer{}
	l.Source(strings.NewReader(input))

	for i, exp := range expected {
		got := l.NextToken()
		if exp != got.Type {
			t.Errorf("Expected token type %d, got %d for token %d", exp, got.Type, i)
		}
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
			res := l.NextToken()
			if res.Type != tokens.NUMBER {
				t.Errorf("Token type incorrect, wanted number, got: %#+v", res)
			}
			if res.Raw != test {
				t.Errorf("Token value incorrect, wanted %q, got: %q", test, res.Raw)
			}
		})
	}
}
