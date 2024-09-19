package script

import (
	"mvdan.cc/sh/v3/syntax"
)

func getSourceFiles(node syntax.Node) (*[]string, error) {
	rt := []string{}
	var err error = nil

	syntax.Walk(node, func(n syntax.Node) bool {
		switch nt := n.(type) {
		case *syntax.Stmt:
			nodes := getCmdCallExprs(nt, "source")
			for _, node := range nodes {
				for _, arg := range node.Args[1:] {
					rt = append(rt, getString(arg.Parts[0]))
				}
			}
		}
		return true
	})

	return &rt, err
}
