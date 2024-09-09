package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
		return
	}
	log.SetFlags(0)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
}
