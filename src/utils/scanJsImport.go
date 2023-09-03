package utils

import "strings"

func ScanJSImport(text string) string {
	splitedString := strings.Split(text, " ")

	if splitedString[0] == "import" && splitedString[2] == "from" {
		return splitedString[3]
	}

	return ""
}
