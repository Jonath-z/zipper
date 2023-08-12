package main

import "github.com/Jonath-z/zipper/src"

func main() {
	paths := src.FilePathMapping()
	println(len(paths))
}
