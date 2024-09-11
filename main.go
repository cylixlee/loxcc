package main

import (
	"log"
	"loxcc/internal"
	"loxcc/internal/frontend/parser"
	"loxcc/internal/frontend/scanner"
	"os"
)

func main() {
	// // check command-line arguments
	// if len(os.Args) != 2 {
	// 	fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
	// 	return
	// }
	// log.SetFlags(0)

	// // read file
	// b, err := os.ReadFile(os.Args[1])
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// source := string(b)

	// read file
	b, err := os.ReadFile("example/benchmark.lox")
	if err != nil {
		log.Fatalln(err.Error())
	}
	source := string(b)

	tokens, err := scanner.Scan(source)
	if err != nil {
		log.Fatalln(err.Error())
	}

	program, err := parser.Parse(tokens)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, decl := range program {
		internal.Inspect(decl)
	}
}
