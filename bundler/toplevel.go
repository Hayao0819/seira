package bundler

import (
	"github.com/Hayao0819/seira/script"
	"github.com/samber/lo"
	"mvdan.cc/sh/v3/syntax"
)

func hasMainFunction(i *script.Script) bool {
	mainFunc := lo.Filter(i.TopLevelFuncs, func(fn syntax.FuncDecl, index int) bool {
		return fn.Name.Value == "main"
	})

	return len(mainFunc) > 0

}
