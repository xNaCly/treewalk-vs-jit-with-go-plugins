package eval

import (
	"errors"
	"fmt"
	"interpreter/ast"
	"interpreter/tokens"
	"os"
	"os/exec"
	"plugin"
	"strings"
)

func generateGo(b *strings.Builder, node ast.Node) {
	switch node := node.(type) {
	case *ast.Binary:
		generateGo(b, node.Left)
		switch node.Token.Type {
		case tokens.PLUS:
			b.WriteRune('+')
		case tokens.MINUS:
			b.WriteRune('-')
		case tokens.ASTERIKS:
			b.WriteRune('*')
		case tokens.SLASH:
			b.WriteRune('/')
		}
		generateGo(b, node.Right)
	case *ast.Unary:
		b.WriteRune('(')
		generateGo(b, node.Right)
		b.WriteString("*-1)")
	case *ast.Number:
		b.WriteString(node.String())
	}
}

func JIT(node ast.Node) (func() float64, error) {
	b := &strings.Builder{}
	b.WriteString("package main\nfunc Main()float64{return ")
	generateGo(b, node)
	b.WriteRune('}')

	dir, err := os.MkdirTemp("", "gojit")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, "jit_*.go")
	if err != nil {
		return nil, err
	}

	file.WriteString(b.String())

	sharedObjectPath := strings.TrimSuffix(file.Name(), ".go")

	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", sharedObjectPath, file.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out), cmd.Args)
		return nil, err
	}

	plug, err := plugin.Open(sharedObjectPath)
	if err != nil {
		return nil, err
	}

	symbol, err := plug.Lookup("Main")
	if err != nil {
		fmt.Println(sharedObjectPath)
		return nil, err
	}

	Main, ok := symbol.(func() float64)
	if !ok {
		return nil, errors.New("Error while accessing jit compiled symbols")
	}

	return Main, nil
}
