package bundler

import (
	"errors"
	"io"

	"mvdan.cc/sh/v3/syntax"
)

type ImportFile struct {
	FilePath string
	Alias    string
}

func getImportInfo(node syntax.Node) (*[]ImportFile, error) {
	rt := []ImportFile{}
	var err error = nil

	syntax.Walk(node, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.Stmt:
			nodes := getCmdCallExprs(nt, "import")
			for _, node := range nodes {
				switch len(node.Args) {
				case 1, 3:
					err = errors.New("not enough arguments for import")
					return false
				case 2:
					src := getString(node.Args[1])

					rt = append(rt, ImportFile{
						FilePath: src,
						Alias:    "",
					})
				case 4:
					src := getString(node.Args[1])
					as := getString(node.Args[2])
					alias := getString(node.Args[3])

					if as != "as" {
						err = errors.New("invalid import syntax")
						return false
					}
					rt = append(rt, ImportFile{
						FilePath: src,
						Alias:    alias,
					})

				}
			}
		}
		return true
	})

	return &rt, err
}

type Script struct {
	Imports       []ImportFile
	TopLevelFuncs []syntax.FuncDecl
	File          *syntax.File
}

func getScriptInfo(r io.Reader, name string) (*Script, error) {
	rt := Script{}
	var err error = nil
	parser := parserWithoutComments
	parsed, err := parser.Parse(r, name)
	if err != nil {
		return nil, err
	}

	rt.File = parsed
	imports, err := getImportInfo(parsed)
	if err != nil {
		return nil, err
	}
	rt.Imports = *imports
	rt.TopLevelFuncs = *getExportedTopLevelFuncList(parsed.Stmts)
	return &rt, err
}

// func getImportInfosFromReader(input io.Reader) (*[]ImportFile, error) {
// 	rt := []ImportFile{}
// 	var err error = nil
// 	walkReader(input, func(s *syntax.Stmt) bool {
// 		srcs, err := getImportInfo(s)
// 		if err != nil {
// 			return false
// 		}
// 		if srcs != nil {
// 			rt = append(rt, *srcs...)
// 		}
// 		return true
// 	})
// 	return &rt, err
// }
