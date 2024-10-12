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
	CC struct {
		Path          string
		Args          []string
		CleanUpSource bool `yaml:"cleanUpSource"`
	}
	Formatter struct {
		Path string
		Args []string
	}
	OutputFolder string `yaml:"outputFolder"`
}

func Build(config BuildConfig, filename, source string) {
	dir := filepath.Join(filepath.Dir(filename), config.OutputFolder)
	deletion := config.CC.Path != "" && config.CC.CleanUpSource

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
	if config.CC.Path != "" {
		outputName := filepath.Join(dir, removeExt(filepath.Base(correspond)))
		cc := newSystemCC(config.CC.Path, outputName)
		args := stl.MakeVectorOf(correspond)
		for source := range unpacker.Sources() {
			args.PushBack(source)
		}
		args.Append(config.CC.Args...)
		cc.Compile(args)
	}

	// (optional) call formatter
	if !deletion && config.Formatter.Path != "" {
		args := stl.MakeVectorOf(correspond)
		args.Append(config.Formatter.Args...)
		command := exec.Command(config.Formatter.Path, args...)
		if err := command.Run(); err != nil {
			log.Fatalln(err.Error())
		}
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
