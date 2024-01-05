package main

import (
	"encoding/json"
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
	debug := flag.Bool("debug", false, "print debug information and traces")
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

	if *debug {
		out, _ := json.MarshalIndent(tree, "", "\t")
		fmt.Println(string(out))
	}

	for tree != nil {
		var out float64
		switch {
		case *ast:
			out = eval.TreeWalk(tree)
		case *vm:
			vm := eval.Vm{}
			instructions := eval.BytecodeCompile(&vm, tree)
			if *debug {
				for _, op := range eval.InspectBytecode(instructions) {
					fmt.Println(op)
				}
			}
			out = vm.Eval(instructions)
		case *jit:
			panic("Not implemented")
		default:
			log.Fatalln("No evaluation strategy supplied")
		}
		fmt.Printf("%f\n", out)
		tree = p.Next()
	}
}
