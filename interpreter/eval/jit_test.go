package eval

import (
	"interpreter/lexer"
	"interpreter/parser"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJIT(t *testing.T) {
	tests := []struct {
		input  string
		output float64
	}{
		{"1+2", 3},
		{"1-2", -1},
		{"1/2", 0.5},
		{"1*2", 2},
		{"1+2+3", 6},
		{"-3", -3},
		{"--3", 3},
	}

	for _, test := range tests {
		l := lexer.Lexer{}
		l.Source(strings.NewReader(test.input))
		p := parser.Parser{}
		p.Source(l)
		tree := p.Next()
		function, error := JIT(tree)
		out := function()
		assert.NoError(t, error)
		assert.Equal(t, test.output, out)
	}
}
