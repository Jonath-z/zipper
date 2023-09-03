package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type NodeDependency struct {
	line int64
	path string
}
type Node struct {
	path         string
	content      string
	dependencies []*NodeDependency
}

func MapNodeDepencies(path string) *[]NodeDependency {
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
		counter++
		// Continue from here
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scanning the file")
	}

}

func CreateNode(path string) *Node {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Can not read the file on %v", path)
		return nil
	}

	//file, err := parser.ParseFile(nil, path, string(fileContent), 1)

	//println(string(fileContent))

	MapNodeDepencies(path)

	node := &Node{
		path:         path,
		content:      string(fileContent),
		dependencies: nil,
	}

	return node
}
