package main

import (
	"log"
	"loxcc/internal"
	"loxcc/internal/frontend/parser"
	"loxcc/internal/frontend/scanner"
	"os"

	stl "github.com/chen3feng/stl4go"
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
	b, err := os.ReadFile("example/expr.lox")
	if err != nil {
		log.Fatalln(err.Error())
	}
	source := string(b)

	// scan source
	var tokens stl.Vector[*scanner.Token]
	s := scanner.NewScanner(source)
	for {
		token, err := s.Scan()
		if err != nil {
			log.Fatalln(err.Error())
		}

		if token == nil {
			break
		}
		tokens.PushBack(token)
	}

	p := parser.NewParser(tokens)
	expr, err := p.ParseExpression()
	if err != nil {
		log.Fatal(err.Error())
	}

	internal.Inspect(expr)
}
