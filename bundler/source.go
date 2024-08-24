package bundler

import (
	"io"

	"github.com/Hayao0819/seira/utils"
	"mvdan.cc/sh/v3/syntax"
)

// TODO: Implement this function
func getImportedFilePaths(src io.Reader) *[]string {
	parser := syntax.NewParser(syntax.KeepComments(true))

	parser.Stmts(src, func(s *syntax.Stmt) bool {
		// TODO: import 文のみを抽出
		utils.PrintAsJSON(s)
		return true
	})
	return nil
}
