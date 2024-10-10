package internal

import (
	"log"
	"loxcc/assets"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	stl "github.com/chen3feng/stl4go"
)

type systemCC struct {
	path       string
	outputName string
}

func newSystemCC(path, outputName string) systemCC {
	return systemCC{
		path:       path,
		outputName: outputName,
	}
}

func (sc systemCC) Compile(sources stl.Vector[string]) {
	// command arguments
	args := stl.MakeVectorOf("-o", sc.outputName)
	args.Append(sources...)

	cmd := exec.Command(sc.path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

type BuildConfig struct {
	OutputFolderName       string `yaml:"outputFolderName"`
	CcPath                 string `yaml:"ccPath"`
	DeleteSourceAfterBuild bool   `yaml:"deleteSourceAfterBuild"`
}

func Build(config BuildConfig, filename, source string) {
	dir := filepath.Join(filepath.Dir(filename), config.OutputFolderName)
	deletion := config.CcPath != "" && config.DeleteSourceAfterBuild

	// unpack LOXCRT
	unpacker := assets.NewRuntimeUnpacker(dir)
	unpacker.Unpack()
	if deletion {
		defer unpacker.Remove()
	}

	// output generated C code
	correspond := filepath.Join(dir, filepath.Base(filename)+".c")
	if err := os.WriteFile(correspond, []byte(source), 0666); err != nil {
		log.Fatalln(err.Error())
	}
	if deletion {
		defer os.Remove(correspond)
	}

	// (optional) call system CC
	if config.CcPath != "" {
		outputName := filepath.Join(dir, removeExt(filepath.Base(correspond)))
		cc := newSystemCC(config.CcPath, outputName)
		sources := stl.MakeVectorOf(correspond)
		for source := range unpacker.Sources() {
			sources.PushBack(source)
		}
		cc.Compile(sources)
	}
}

func removeExt(path string) string {
	for {
		ext := filepath.Ext(path)
		if ext == "" {
			return path
		}
		path = strings.TrimSuffix(path, ext)
	}
}
