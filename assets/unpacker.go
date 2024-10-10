package assets

import (
	"io/fs"
	"iter"
	"log"
	"os"
	"path/filepath"
	"strings"

	stl "github.com/chen3feng/stl4go"
)

// The extensions that distinguish source files from others.
//
// For RuntimeUnpacker, it provides a way to iterate over those source files, in order to
// simplify the logic of system CC invocation.
//
// The extensions are case-insensitive because they're all ToLower'ed in
// RuntimeUnpacker.Sources().
var sourceExts = stl.MakeBuiltinSetOf(".c")

// The universal manager of LOXCRT files.
//
// Before compilation through the system CC, LOXCRT implementation files should be copied
// to the output directory of the Lox program; after that, they may be deleted according
// to the build-config.
//
// RuntimeUnpacker is responsible for managing the LOXCRT files: it unpacks them to the
// target path, and holds their paths internally, which is very useful when deleting them
// later. It also provides a way to iterate over the source files (.c), in order to
// simplify the logic of system CC invocation.
type RuntimeUnpacker struct {
	path     string
	unpacked *stl.DList[string]
}

// Create a RuntimeUnpacker, specifying the output path.
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

// Unpacks the LOXCRT.
//
// Since the LOXCRT files are embedded as embed.FS, this function internally calls
// fs.WalkDir to iterate over the files and directories, to copy them to the specified
// path.
//
// The copy routine preserves the structure of embedded filesystem; it creates the
// corresponding file or directory if there does not exist one. If the file already
// exists, it overwrites them to ensure the consistency of LOXCRT.
//
// Files are created under 0666 permission, while the directories are under 0777.
//
// In this procedure, the real paths of copied files and directories are recorded
// internally, which is for possible future removal operation.
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

// Removes the unpacked files from OS filesystem.
//
// It deletes all files and directories unpacked previously. If one file or directory is
// already deleted, it will do nothing.
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

// Iterator over all source files.
//
// More specifically, since we're transpiling to C, source files refer to the files with
// .c extension. Header files are not compilation units in C and we don't need to pass
// them to the system CC.
//
// Note: make sure you're compiling this on Go 1.23.2 or later, or you may get some
// internal bugs due to inadequate implementation of Go range over func.
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
