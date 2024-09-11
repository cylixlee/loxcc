package assets

import (
	"embed"
	"text/template"
)

var (
	//go:embed main.tpl
	tpl string

	//go:embed runtime
	Runtime    embed.FS
	Entrypoint *template.Template
)

func init() {
	t, err := template.New("main").Parse(tpl)
	if err != nil {
		panic(err)
	}
	Entrypoint = t
}
