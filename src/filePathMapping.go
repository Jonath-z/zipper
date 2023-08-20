package src

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FilePathMapping() []string {
	var jsFilesPath []string

	e := filepath.Walk("js", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasSuffix(path, ".js") {
			jsFilesPath = append(jsFilesPath, path)
			println(path)
		}
		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	return jsFilesPath
}
