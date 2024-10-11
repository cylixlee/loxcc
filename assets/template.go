package assets

import (
	"embed"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	// The text templates to create a C source file according to a Lox program.
	//
	// Different C declarations (e.g. functions, variables) should be placed at different
	// segments of a single C source file, and each of them have a different syntax, so
	// multiple text template is introduced.
	//
	// The template files are embedded for convenience.
	//
	//go:embed template
	fsys embed.FS

	// Parsed templates according to embedded template files.
	Templates = template.New("")
)

// parse the templates recursively when package is imported
func init() {
	var err error

	Templates.Funcs(map[string]any{
		"minus": func(a, b int) int {
			return a - b
		},
	})
	_, err = Templates.ParseFS(fsys, "template/*.tpl")
	if err != nil {
		panic(err)
	}
}

// Shorthand for calling template.ExecuteTemplate.
func ApplyTemplate(name string, data any) string {
	var builder strings.Builder

	if err := Templates.ExecuteTemplate(&builder, name, data); err != nil {
		panic(err)
	}
	return builder.String()
}

func removeExt(path string) string {
	for {
		ext := filepath.Ext(path)
		if ext == "" {
			break
		}
		path = strings.TrimSuffix(path, ext)
	}
	return path
}
