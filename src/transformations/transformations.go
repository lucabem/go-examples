package Transformations

import "strings"

func ReplaceExtension(input string) string {
	return strings.ReplaceAll(input, "xlsx", "csv")
}

func ReplaceSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "_")
}

func ToLower(input string) string {
	return strings.ToLower(input)
}

func Trim(input string) string {
	return strings.Trim(input, "")
}
