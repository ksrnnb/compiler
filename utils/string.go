package utils

import "strings"

func TrimSpaceAndTab(text string) string {
	trimmedText := strings.Trim(text, "\t")
	trimmedText = strings.Trim(trimmedText, " ")

	if len(trimmedText) == 0 {
		return ""
	}

	if trimmedText[0] == '\t' {
		return TrimSpaceAndTab(trimmedText)
	}

	return trimmedText
}
