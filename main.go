package main

import (
	"fmt"
	"log"
	"loxcc/internal/frontend/scanner"
	"os"
)

func main() {
	// check command-line arguments
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
		return
	}
	log.SetFlags(0)

	// read file
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}
	source := string(b)

	// scan source
	s := scanner.NewScanner(source)
	for {
		token, err := s.Scan()
		if err != nil {
			log.Fatalln(err.Error())
		}

		if token == nil {
			break
		}
		fmt.Printf("%4d %2d %s\n", token.Lineno, token.Type, token.Lexeme)
	}
}
