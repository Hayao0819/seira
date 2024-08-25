package bundler

import (
	"io"
	"strings"

	"github.com/samber/lo"
	"mvdan.cc/sh/v3/syntax"
)

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

func getImportedFilePaths(node syntax.Node) *[]string {
	rt := []string{}

	syntax.Walk(node, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.Stmt:
			nodes := getCmdCallExprs(nt, "import")
			for _, node := range nodes {
				for _, arg := range node.Args[1:] {
					rt = append(rt, getString(arg))
				}
			}
		}
		return true
	})

	return &rt
}

func getImportedFilePathsFromReader(input io.Reader) *[]string {
	rt := []string{}

	walkReader(input, func(s *syntax.Stmt) bool {
		srcs := getImportedFilePaths(s)
		if srcs != nil {
			lo.ForEach(*srcs, func(src string, index int) {
				rt = append(rt, src)
			})
		}
		return true
	})
	return &rt
}

func walkReader(input io.Reader, f func(*syntax.Stmt) bool) {
	parser := parserWithComments
	parser.Stmts(input, func(s *syntax.Stmt) bool {
		return f(s)
	})
}
