package script

import (
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

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

// func walkReader(input io.Reader, f func(*syntax.Stmt) bool) {
// 	parser := parserWithComments
// 	parser.Stmts(input, func(s *syntax.Stmt) bool {
// 		return f(s)
// 	})
// }
