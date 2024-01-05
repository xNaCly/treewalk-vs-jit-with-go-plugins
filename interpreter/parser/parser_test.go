package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/tokens"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input  string
		result ast.Node
	}{
		{
			input: "1+2",
			result: &ast.Binary{
				Token: tokens.Token{
					Type:   tokens.PLUS,
					Line:   1,
					Column: 2,
				},
				Left: &ast.Number{
					Token: tokens.Token{
						Raw:    "1",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 1,
					},
					Value: 1,
				},
				Right: &ast.Number{
					Token: tokens.Token{
						Raw:    "2",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 2,
					},
					Value: 2,
				},
			},
		},
		{
			input: "1*2",
			result: &ast.Binary{
				Token: tokens.Token{
					Type:   tokens.ASTERIKS,
					Line:   1,
					Column: 2,
				},
				Left: &ast.Number{
					Token: tokens.Token{
						Raw:    "1",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 1,
					},
					Value: 1,
				},
				Right: &ast.Number{
					Token: tokens.Token{
						Raw:    "2",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 2,
					},
					Value: 2,
				},
			},
		},
		{
			input: "1/2",
			result: &ast.Binary{
				Token: tokens.Token{
					Type:   tokens.SLASH,
					Line:   1,
					Column: 2,
				},
				Left: &ast.Number{
					Token: tokens.Token{
						Raw:    "1",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 1,
					},
					Value: 1,
				},
				Right: &ast.Number{
					Token: tokens.Token{
						Raw:    "2",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 2,
					},
					Value: 2,
				},
			},
		},
		{
			input: "1-2",
			result: &ast.Binary{
				Token: tokens.Token{
					Type:   tokens.MINUS,
					Line:   1,
					Column: 2,
				},
				Left: &ast.Number{
					Token: tokens.Token{
						Raw:    "1",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 1,
					},
					Value: 1,
				},
				Right: &ast.Number{
					Token: tokens.Token{
						Raw:    "2",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 2,
					},
					Value: 2,
				},
			},
		},
		{
			input: "1-2*5",
			result: &ast.Binary{
				Token: tokens.Token{
					Type:   tokens.MINUS,
					Line:   1,
					Column: 2,
				},
				Left: &ast.Number{
					Token: tokens.Token{
						Raw:    "1",
						Type:   tokens.NUMBER,
						Line:   1,
						Column: 1,
					},
					Value: 1,
				},
				Right: &ast.Binary{
					Token: tokens.Token{
						Type:   tokens.ASTERIKS,
						Line:   1,
						Column: 4,
					},
					Left: &ast.Number{
						Token: tokens.Token{
							Raw:    "2",
							Type:   tokens.NUMBER,
							Line:   1,
							Column: 3,
						},
						Value: 2,
					},
					Right: &ast.Number{
						Token: tokens.Token{
							Raw:    "5",
							Type:   tokens.NUMBER,
							Line:   1,
							Column: 4,
						},
						Value: 5,
					},
				},
			},
		},
	}

	for _, test := range tests {
		l := lexer.Lexer{}
		l.Source(strings.NewReader(test.input))
		p := Parser{}
		p.Source(l)
		out := p.Next()
		assert.EqualValues(t, test.result, out)
	}
}
