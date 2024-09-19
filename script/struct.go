package script

import "mvdan.cc/sh/v3/syntax"

type Script struct {
	Imports       *[]string
	TopLevelFuncs []syntax.FuncDecl
	File          *syntax.File
	FullPath      string
}
