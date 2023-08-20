package main

import (
	"fmt"
	"github.com/Jonath-z/zipper/src"
	"github.com/dop251/goja/parser"
	"log"
	"os"
)

func main() {
	path := src.FilePathMapping()
	content, err := os.ReadFile(path[0])
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, err := parser.ParseFile(nil, path[0], string(content), 1)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(file.File)
}
