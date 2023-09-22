package utils

import (
	"bufio"
	"fmt"
	"os"
)

func FileScanner(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
