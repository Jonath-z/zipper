package src

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FileMapping() []string {
	var fileMap []string

	e := filepath.Walk("js", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasSuffix(path, ".js") {
			fileMap = append(fileMap, path)
		}

		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(fileMap)
	return fileMap
}
