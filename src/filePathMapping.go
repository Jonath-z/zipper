package src

import (
	"log"
	"os"
	"path/filepath"
)

func FilePathMapping() []string {
	var jsFilesPath []string

	e := filepath.Walk("js", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		jsFilesPath = append(jsFilesPath, path)
		println(info)
		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	return jsFilesPath
}
