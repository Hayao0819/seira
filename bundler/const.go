package bundler

import "mvdan.cc/sh/v3/syntax"

var parserWithComments = syntax.NewParser(syntax.KeepComments(true))
var parserWithoutComments = syntax.NewParser()
