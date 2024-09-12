package main

import (
	"fmt"
	"log"
	"loxcc/assets"
	"loxcc/internal/backend"
	"loxcc/internal/frontend/parser"
	"loxcc/internal/frontend/scanner"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	stl "github.com/chen3feng/stl4go"
	"gopkg.in/yaml.v3"
)

const (
	ConfigPath  = "build-config.yaml"
	RuntimePath = "runtime"
)

var config = buildConfig{
	OutputFolderName:       "build",
	CcPath:                 "",
	DeleteSourceAfterBuild: false,
}

type buildConfig struct {
	OutputFolderName       string `yaml:"outputFolderName"`
	CcPath                 string `yaml:"ccPath"`
	DeleteSourceAfterBuild bool   `yaml:"deleteSourceAfterBuild"`
}

func init() {
	log.SetFlags(0)

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return // if there's no config file, use the default value instead of panic.
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: loxcc [path]")
		return
	}

	outputFolder := filepath.Join(filepath.Dir(os.Args[1]), config.OutputFolderName)
	prepareFolder(outputFolder)

	code := compile(os.Args[1])
	writeMain(outputFolder, code)

	paths := copyRuntime(outputFolder)
	if config.CcPath != "" {
		paths.PushBack(filepath.Join(outputFolder, "main.c"))
		paths.Append("-o", filepath.Join(outputFolder, "a.exe"))
		cmd := exec.Command(config.CcPath, paths...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func compile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	source := string(data)

	tokens, err := scanner.Scan(source)
	if err != nil {
		log.Fatalln(err.Error())
	}

	program, err := parser.Parse(tokens)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return backend.Generate(program)
}

func prepareFolder(folder string) {
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		log.Fatalln(err.Error())
	}
}

func copyRuntime(folder string) stl.Vector[string] {
	var paths stl.Vector[string]

	entries, err := assets.Runtime.ReadDir(RuntimePath)
	if err != nil {
		panic(err)
	}

	for _, v := range entries {
		// read runtime sources from embed.FS
		data, err := assets.Runtime.ReadFile(path.Join(RuntimePath, v.Name()))
		if err != nil {
			panic(err)
		}

		// write to output folder (with rw- permission)
		p := filepath.Join(folder, v.Name())
		err = os.WriteFile(p, data, 0666)
		if err != nil {
			log.Fatalln(err.Error())
		}
		paths.PushBack(p)
	}
	return paths
}

func writeMain(folder, source string) {
	file, err := os.Create(filepath.Join(folder, "main.c"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	if _, err := file.WriteString(source); err != nil {
		log.Fatalln(err.Error())
	}
}
