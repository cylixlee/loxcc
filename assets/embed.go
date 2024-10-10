package assets

import (
	"embed"
	"text/template"
)

var (
	// The text template to create a C source file according to a Lox program.
	//
	// Different C declarations (e.g. functions, variables) should be placed at different
	// segments of a single C source file, so a text template is introduced.
	//
	// The template file is embedded for convenience.
	//
	//go:embed main.tpl
	tpl string

	// The Lox C Runtime.
	//
	// To provide the dynamic features and extra language components (e.g. GC), some
	// runtime preparation is needed. Since we're transpiling Lox to C, the runtime part
	// called LOXCRT is implemented in C.
	//
	// Before compilation, the LOXCRT files are copied to the output directory, and
	// compiled together with the template-generated C code into an executable.
	//
	// The whole directory containing all the runtime implementation is embedded.
	//
	//go:embed runtime
	rt embed.FS

	// The text template of generated C code.
	//
	// For now, each Lox source file will generate exactly one C file. Different syntax
	// elements (e.g. functions, variables) are transpiled as corresponding C code at
	// different segments of the generated file.
	//
	// This template is parsed when the packaged is imported, from the embedded template
	// text.
	Entrypoint *template.Template
)

func init() {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	Entrypoint = t
}
