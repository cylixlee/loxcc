package main

import (
	"fmt"
	"log"
	"loxcc/internal"
	"loxcc/internal/analyzer"
	"loxcc/internal/codegen"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
		return
	}

	// define default build-config, and load from file if exists.
	config := internal.BuildConfig{OutputFolder: "build"}
	data, err := os.ReadFile("build-config.yaml")
	if err == nil { // will not report an error if ReadFile fails.
		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatalln(err.Error())
		}
	}

	data, err = os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err.Error())
	}

	program, err := analyzer.Analyze(string(data))
	if err != nil {
		log.Fatalln(err.Error())
	}
	code := codegen.Generate(program)

	internal.Build(config, os.Args[1], code)
}
