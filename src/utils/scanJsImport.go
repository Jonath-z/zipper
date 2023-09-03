package utils

import "strings"

func ScanJSImport(text string) string {
	splittedString := strings.Split(text, " ")

	if splittedString[0] == "import" && splittedString[2] == "from" {
		return splittedString[3]
	}

	return ""
}
