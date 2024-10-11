package assets

import (
	_ "embed"
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
