package assets

import (
	"embed"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	stl "github.com/chen3feng/stl4go"
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

func CopyRuntime(folder string) stl.Vector[string] {
	filenames := stl.MakeVector[string]()

	runtimePath := filepath.Join(folder, "runtime")
	if err := os.MkdirAll(runtimePath, 0666); err != nil {
		log.Fatalln(err.Error())
	}

	entries, err := Runtime.ReadDir("runtime")
	if err != nil {
		panic(err)
	}

	for _, v := range entries {
		// reading from embed.FS should not fail, program panics if so
		from := path.Join("runtime", v.Name())
		data, err := Runtime.ReadFile(from)
		if err != nil {
			panic(err)
		}

		// writing to OS filesystem may fail, message is logged if so
		to := filepath.Join(runtimePath, v.Name())
		if err := os.WriteFile(to, data, 0666); err != nil {
			log.Fatalln(err.Error())
		}

		filenames.PushBack(to)
	}
	return filenames
}
