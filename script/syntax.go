package script

import (
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

// func getDefinedFuncList(node syntax.Node) []string {
// 	rt := []string{}
// 	syntax.Walk(node, func(n syntax.Node) bool {
// 		switch nt := n.(type) {
// 		case *syntax.FuncDecl:
// 			rt = append(rt, nt.Name.Value)
// 		}
// 		return true
// 	})
// 	return rt
// }

func getExportedTopLevelFuncList(stmts []*syntax.Stmt) *[]syntax.FuncDecl {
	rt := []syntax.FuncDecl{}
	for _, stmt := range stmts {
		switch stmt.Cmd.(type) {
		case *syntax.FuncDecl:
			rt = append(rt, *stmt.Cmd.(*syntax.FuncDecl))
		}
	}
	return &rt
}
