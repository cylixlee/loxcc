package assets

import (
	"embed"
	"io/fs"
	"iter"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	stl "github.com/chen3feng/stl4go"
)

var (
	//go:embed main.tpl
	tpl string
	//go:embed runtime
	rt embed.FS

	sourceExts = stl.MakeBuiltinSetOf(".c")
	Entrypoint *template.Template
)

func init() {
	t, err := template.New("main").Parse(tpl)
	if err != nil {
		panic(err)
	}
	Entrypoint = t
}

type RuntimeUnpacker struct {
	path     string
	unpacked *stl.DList[string]
}

func NewRuntimeUnpacker(path string) *RuntimeUnpacker {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &RuntimeUnpacker{
		path:     path,
		unpacked: stl.NewDList[string](),
	}
}

func (ru *RuntimeUnpacker) Unpack() {
	err := fs.WalkDir(rt, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." { // do nothing at root
			return nil
		}
		correspond := filepath.Join(ru.path, path)

		if d.IsDir() {
			// mkdir if the entry is dir
			if err := os.MkdirAll(correspond, 0777); err != nil {
				return err
			}
		} else {
			// read data from file
			data, err := rt.ReadFile(path)
			if err != nil {
				return err
			}

			// write to OS filesystem
			if err := os.WriteFile(correspond, data, 0666); err != nil {
				return err
			}
		}
		ru.unpacked.PushBack(correspond)
		return nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (ru *RuntimeUnpacker) Remove() {
	for {
		path, exist := ru.unpacked.PopBack()
		if !exist {
			break
		}
		if err := os.RemoveAll(path); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func (ru RuntimeUnpacker) Sources() iter.Seq[string] {
	return func(yield func(string) bool) {
		ru.unpacked.ForEachIf(func(path string) bool {
			ext := strings.ToLower(filepath.Ext(path))
			if sourceExts.Has(ext) {
				if !yield(path) {
					return false
				}
			}
			return true
		})
	}
}
