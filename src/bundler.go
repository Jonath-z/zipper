package src

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Bundler() {
	outputFile := "/src/js/dist/index.js"
	fileMap := FileMapping()

	//dependencyGraph := CreateDependencyGraph()

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

	for _, path := range fileMap {
		fmt.Println(path)
	}
}
