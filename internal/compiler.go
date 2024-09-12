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

type BuildConfig struct {
	OutputFolderName       string `yaml:"outputFolderName"`
	CcPath                 string `yaml:"ccPath"`
	DeleteSourceAfterBuild bool   `yaml:"deleteSourceAfterBuild"`
}

func Compile(config BuildConfig, path, source string) {
	parentFolder := filepath.Dir(path)
	outputFolder := filepath.Join(parentFolder, config.OutputFolderName)
	ext := filepath.Ext(path)
	filename := filepath.Base(path)
	filenameWithoutSuffix := strings.TrimSuffix(filename, ext)

	// prepare output folder
	if err := os.MkdirAll(outputFolder, 0666); err != nil {
		log.Fatalln(err.Error())
	}

	// write generated source to file.
	sourcePath := filepath.Join(outputFolder, filename) + ".c"
	if err := os.WriteFile(sourcePath, []byte(source), 0666); err != nil {
		log.Fatalln(err.Error())
	}

	// copy runtime
	rt := assets.CopyRuntime(outputFolder)

	// (optional) call CC
	if config.CcPath != "" {
		binaryPath := filepath.Join(outputFolder, filenameWithoutSuffix)
		args := stl.MakeVectorOf("-o", binaryPath, sourcePath)
		args.Append(rt...)

		cmd := exec.Command(config.CcPath, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalln(err.Error())
		}

		// (optional) cleanup C source
		if config.DeleteSourceAfterBuild {
			for _, v := range rt {
				if err := os.Remove(v); err != nil {
					log.Fatalln(err.Error())
				}
			}
			if err := os.Remove(sourcePath); err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
