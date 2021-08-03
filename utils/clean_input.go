package utils

import "strings"

func CleanInput(input string) string {
	// Replace CRLF in the input string
	input = strings.Replace(input, "\r\n", "", -1)
	// Replace LF in the input string
	input = strings.Replace(input, "\n", "", -1)
	return input
}
