// package main

// import (
// 	"fmt"
// 	"log"
// 	"loxcc/internal"
// 	"loxcc/internal/analyzer"
// 	"loxcc/internal/backend"
// 	"os"

// 	"gopkg.in/yaml.v3"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
// 		return
// 	}

// 	// define default build-config, and load from file if exists.
// 	config := internal.BuildConfig{
// 		OutputFolderName:       "build",
// 		CcPath:                 "",
// 		DeleteSourceAfterBuild: false,
// 	}
// 	data, err := os.ReadFile("build-config.yaml")
// 	if err == nil { // will not report an error if ReadFile fails.
// 		if err := yaml.Unmarshal(data, &config); err != nil {
// 			log.Fatalln(err.Error())
// 		}
// 	}

// 	// AST codegen
// 	data, err = os.ReadFile(os.Args[1])
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	program, err := analyzer.Analyze(string(data))
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	code := backend.Generate(program)

// 	// internal.Compile(config, os.Args[1], code)
// 	internal.Build(config, os.Args[1], code)
// }

package main

import (
	"fmt"
	"loxcc/assets"
)

type NumberType float64

func main() {
	fmt.Println(assets.Templates.DefinedTemplates())
	fmt.Println(assets.ApplyTemplate("nil", nil))
	fmt.Println(assets.ApplyTemplate("boolean", true))

	var number NumberType = 12.34
	fmt.Println(assets.ApplyTemplate("number", number))
	fmt.Println(assets.ApplyTemplate("string", "hello"))
	fmt.Println(assets.ApplyTemplate("binary", map[string]string{
		"left":         "1",
		"right":        "2",
		"operatorFunc": "LRT_Add",
	}))
	fmt.Println(assets.ApplyTemplate("unary", map[string]string{
		"operand":      "NIL",
		"operatorFunc": "LRT_Not",
	}))
}
