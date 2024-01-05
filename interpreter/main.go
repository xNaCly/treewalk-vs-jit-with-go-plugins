package main

import (
	"flag"
	"fmt"
	"interpreter/eval"
	"interpreter/lexer"
	"interpreter/parser"
	"log"
	"os"
)

func main() {
	ast := flag.Bool("ast", false, "use tree-walk evaluation")
	vm := flag.Bool("vm", false, "use bytecode compilation and evaluation")
	jit := flag.Bool("jit", false, "use jit compilation and evaluation")
	inPath := flag.String("i", "", "file to interpret")
	flag.Parse()

	if len(*inPath) == 0 {
		log.Fatalln("No input provided")
	}

	file, err := os.Open(*inPath)
	if err != nil {
		log.Fatalln(err)
	}

	l := lexer.Lexer{}
	l.Source(file)

	p := parser.Parser{}
	p.Source(l)
	tree := p.Next()

	for tree != nil {
		var out float64
		switch {
		case *ast:
			out = eval.TreeWalk(tree)
		case *jit:
			log.Println("Eval with jit compiler")
		case *vm:
			log.Println("Eval with bytecode compiler and vm")
		default:
			log.Fatalln("No evaluation strategy supplied")
		}
		fmt.Printf("%f\n", out)
		tree = p.Next()
	}
}
