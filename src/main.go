package main

import (
	"flag"
	"log"
)

func main() {
	treeWalk := flag.Bool("ast", false, "instruct interpreter to use the tree-walk backend")
	bytecode := flag.Bool("bytecode", false, "instruct interpreter to use the bytecode backend")
	golang := flag.Bool("jit", false, "instruct interpreter to use the go code generation backend")
	flag.Parse()

	if *treeWalk {
		log.Println("using the tree walk interpreter")
	} else if *bytecode {
		log.Println("using the bytecode interpreter")
	} else if *golang {
		log.Println("using the go JIT")
	} else {
		log.Fatalln("No evaluation backend provided")
	}
}
