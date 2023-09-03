package src

import (
	"bufio"
	"fmt"
	"github.com/Jonath-z/zipper/src/utils"
	"log"
	"os"
)

type NodeDependency struct {
	Line           int64  `json:"line"`
	DependencyPath string `json:"dependencyPath"`
}
type Node struct {
	Path         string           `json:"path"`
	Content      string           `json:"content"`
	Dependencies []NodeDependency `json:"dependencies"`
}

func mapNodeDependencies(path string) []NodeDependency {
	var nodeDependency []NodeDependency

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error while closing the file %v\n", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var counter int64
	for scanner.Scan() {
		if utils.ScanJSImport(scanner.Text()) != "" {
			counter++
			node := &NodeDependency{
				DependencyPath: utils.ScanJSImport(scanner.Text()),
				Line:           counter,
			}

			nodeDependency = append(nodeDependency, *node)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scanning the file")
	}

	return nodeDependency
}

func CreateNode(path string) *Node {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Can not read the file on %v", path)
		return nil
	}

	dependencies := mapNodeDependencies(path)
	node := &Node{
		Path:         path,
		Content:      string(fileContent),
		Dependencies: dependencies,
	}

	return node
}
