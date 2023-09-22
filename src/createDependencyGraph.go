package src

import (
	"encoding/json"
	"fmt"
	"github.com/Jonath-z/zipper/src/utils"
	"os"
	"path/filepath"
	"strings"
)

func CreateDependencyGraph() []Node {
	var dependencyGraph []Node

	e := filepath.Walk("js", func(path string, info os.FileInfo, err error) error {
		utils.CheckError(err)
		if strings.HasSuffix(path, ".js") {
			node := CreateNode(path)
			dependencyGraph = append(dependencyGraph, *node)
		}
		return nil
	})
	utils.CheckError(e)

	b, err := json.Marshal(&dependencyGraph)
	utils.CheckError(err)
	fmt.Println(string(b))
	return dependencyGraph
}
