package src

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateDependencyGraph() []Node {
	var dependencyGraph []Node

	e := filepath.Walk("js", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasSuffix(path, ".js") {
			node := CreateNode(path)
			dependencyGraph = append(dependencyGraph, *node)
			//println(path)
		}

		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	b, err := json.Marshal(&dependencyGraph)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	return dependencyGraph
}
