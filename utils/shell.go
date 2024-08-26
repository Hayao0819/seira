package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

func PrintAsJSON(data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

func Parse(r io.Reader, name string, opts ...syntax.ParserOption) (*syntax.File, error) {
	parser := syntax.NewParser(opts...)
	parsed, err := parser.Parse(r, name)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func ParseFile(path string, opts ...syntax.ParserOption) (*syntax.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	return Parse(f, path, opts...)
}
