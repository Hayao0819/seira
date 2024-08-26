package bundler

import (
	"errors"
	"io"
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

type ImportFile struct {
	FilePath string
	Alias    string
}

func getCmdCallExprs(x syntax.Node, cmd string) []syntax.CallExpr {
	rt := []syntax.CallExpr{}
	syntax.Walk(x, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.CallExpr:
			part := nt.Args[0].Parts[0]
			switch partt := part.(type) {
			case *syntax.Lit:
				if partt.Value == cmd {
					rt = append(rt, *nt)
				}
			}

		}
		return true
	})
	return rt
}

func getString(node syntax.Node) string {
	var sb strings.Builder
	syntax.Walk(node, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.Lit:
			sb.WriteString(nt.Value)
		}
		return true
	})
	return sb.String()
}

func getDefinedFuncList(node syntax.Node) []string {
	rt := []string{}
	syntax.Walk(node, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.FuncDecl:
			rt = append(rt, nt.Name.Value)
		}
		return true
	})
	return rt

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

func getImportInfosFromReader(input io.Reader) (*[]ImportFile, error) {
	rt := []ImportFile{}
	var err error = nil
	walkReader(input, func(s *syntax.Stmt) bool {
		srcs, err := getImportInfo(s)
		if err != nil {
			return false
		}
		if srcs != nil {
			rt = append(rt, *srcs...)
		}
		return true
	})
	return &rt, err
}

func walkReader(input io.Reader, f func(*syntax.Stmt) bool) {
	parser := parserWithComments
	parser.Stmts(input, func(s *syntax.Stmt) bool {
		return f(s)
	})
}
