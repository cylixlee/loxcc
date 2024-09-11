package main

import (
	"fmt"
	"log"
	"loxcc/internal/backend"
	"loxcc/internal/frontend/parser"
	"loxcc/internal/frontend/scanner"
)

func main() {
	source := "(-1 + 2) * 3 - -4;"

	tokens, err := scanner.Scan(source)
	if err != nil {
		log.Fatalln(err.Error())
	}

	program, err := parser.Parse(tokens)
	if err != nil {
		log.Fatalln(err.Error())
	}

	code := backend.Generate(program)
	fmt.Println(code)
}
