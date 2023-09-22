package src

import (
	"fmt"
	"github.com/Jonath-z/zipper/src/utils"
	"strings"
)

type NodeDependency struct {
	Line           int64  `json:"line"`
	DependencyPath string `json:"dependencyPath"`
	Expression     string `json:"expression"`
}
type Node struct {
	Path         string           `json:"path"`
	Content      string           `json:"content"`
	Dependencies []NodeDependency `json:"dependencies"`
}

func getContentWithoutImportAndExportExpression(path string) string {
	var content string
	scanner := utils.FileScanner(path)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "import") || strings.Contains(scanner.Text(), "module.exports") {
			continue
		}
		content += scanner.Text()
	}

	return content
}

func mapNodeDependencies(path string) []NodeDependency {
	var nodeDependency []NodeDependency

	scanner := *utils.FileScanner(path)
	var counter int64
	for scanner.Scan() {
		if utils.ScanJSImport(scanner.Text()) != "" {
			counter++
			node := &NodeDependency{
				DependencyPath: utils.ScanJSImport(scanner.Text()),
				Line:           counter,
				Expression:     utils.ScanModuleImportExpression(scanner.Text()),
			}
			nodeDependency = append(nodeDependency, *node)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error while scanning the file", err)
	}

	return nodeDependency
}

func removeModuleExportLinesFromFileContent(path string) string {
	var content string
	scanner := utils.FileScanner(path)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "module.export") {
			continue
		} else {
			content += scanner.Text()
		}
	}

	println(content)
	return content
}

func CreateNode(path string) *Node {
	dependencies := mapNodeDependencies(path)
	content := removeModuleExportLinesFromFileContent(path)
	//fmt.Println("content without impExp:", getContentWithoutImportAndExportExpression(path))

	node := &Node{
		Path:         path,
		Content:      content,
		Dependencies: dependencies,
	}

	return node
}
