package utils

import "strings"

func FormatName(fileName string) string {
	s := strings.ToLower(fileName)

	if strings.HasSuffix(s, ".txt") {
		return s
	}

	return s + ".txt"
}