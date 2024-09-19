package script

import (
	"io"
	"path/filepath"
)

func GetInfo(r io.Reader, name string) (*Script, error) {
	rt := Script{}
	var err error = nil
	parser := parserWithoutComments
	parsed, err := parser.Parse(r, name)
	if err != nil {
		print("hoge")
		return nil, err
	}

	rt.File = parsed
	sources, err := getSourceFiles(parsed)
	if err != nil {
		return nil, err
	}
	rt.Imports = sources
	rt.TopLevelFuncs = *getExportedTopLevelFuncList(parsed.Stmts)

	abs, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}
	rt.FullPath = abs
	return &rt, nil
}
