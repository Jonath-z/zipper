package src

import (
	"fmt"
	"github.com/Jonath-z/zipper/src/utils"
	"log"
	"os"
	"path/filepath"
)

func Bundler() {
	outputFile := "js/dist/index.js"
	//fileMap := FileMapping()

	dependencyGraph := CreateDependencyGraph()

	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	bundledFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(bundledFile)

	for _, node := range dependencyGraph {
		if node.Path == "js/index.js" {
			content := node.Content
			if node.Dependencies != nil {
				for _, dependency := range node.Dependencies {
					if dependency.Expression != "" {
						fmt.Println("content:", content, "expression:", dependency.Expression)
						//replaced := strings.ReplaceAll(content, dependency.Expression, node.Content)
						//content = replaced
						//fmt.Println("replaced", replaced)
					}
				}
			}
			_, err := bundledFile.Write([]byte(node.Content))
			utils.CheckError(err)
		}
	}

}
