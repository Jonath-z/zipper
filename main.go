package main

import "github.com/Jonath-z/zipper/src"

func main() {
	src.CreateDependencyGraph()
	//path := src.FilePathMapping()
	//content, err := os.ReadFile(path[0])
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}
	//
	//file, err := parser.ParseFile(nil, path[0], string(content), 1)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//fmt.Println(file.File)
}
