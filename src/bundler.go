package src

import (
	"fmt"
	"github.com/Jonath-z/zipper/src/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func findDependencyContent(dependencyPath string, dependencyGraph []Node) string {
	var dependencyContent string
	cleanedPath := strings.Split(dependencyPath, "'")[1]
	cleanedPathWithFileNameOnly := strings.Split(cleanedPath, "/")[1]
	cleanedPathWithPrefix := "js/" + cleanedPathWithFileNameOnly
	fmt.Println("cleanedPath:", cleanedPathWithPrefix)

	for _, node := range dependencyGraph {
		if cleanedPathWithPrefix == node.Path {
			dependencyContent = node.Content
			break
		}
	}

	return dependencyContent
}

func Bundler() {
	outputFile := "js/dist/index.js"

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
						dependencyContent := findDependencyContent(dependency.DependencyPath, dependencyGraph)
						content = dependencyContent + ";" + content
					}
				}
			}
			_, err := bundledFile.Write([]byte(content))
			utils.CheckError(err)
		}
	}

}
