package ast

import (
	stl "github.com/chen3feng/stl4go"
)

// Lox programs consist of declarations.
type Program = stl.Vector[Declaration]
