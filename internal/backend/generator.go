package backend

import (
	"fmt"
	"loxcc/assets"
	"loxcc/internal/ast"
	"strings"

	stl "github.com/chen3feng/stl4go"
)

func Generate(program ast.Program) string {
	var generator codeGenerator
	for _, v := range program {
		v.Accept(&generator)
	}

	return generator.generate()
}

type codeGenerator struct {
	main   stl.Vector[string]
	buffer strings.Builder
}

func (c *codeGenerator) generate() string {
	var builder strings.Builder

	source := strings.Join(c.main, "\n")
	err := assets.Entrypoint.Execute(&builder, source)
	if err != nil {
		panic(err)
	}

	return builder.String()
}

func (c *codeGenerator) write(v ...string) {
	line := strings.Join(v, " ")
	if _, err := c.buffer.WriteString(line); err != nil {
		panic(err)
	}
}

//lint:ignore U1000 currently not used
func (c *codeGenerator) writeln(v ...string) {
	c.write(v...)
	c.write("\n")
}

func (c *codeGenerator) writef(format string, v ...any) {
	line := fmt.Sprintf(format, v...)
	if _, err := c.buffer.WriteString(line); err != nil {
		panic(err)
	}
}

func (c *codeGenerator) push() {
	c.main.PushBack(c.buffer.String())
	c.buffer.Reset()
}
