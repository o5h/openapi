package generator

import (
	"strings"
	"unicode"
)

func trimNonLetterPrefix(s string) string {
	for i, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		return s[i:]
	}
	return ""
}

func toPascalCase(s string) string {
	var result string
	capitalizeNext := true
	for _, r := range s {
		if !isAlphaNumeric(r) {
			capitalizeNext = true
		} else {
			if capitalizeNext {
				result += string(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				result += string(r)
			}
		}
	}
	return result
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func normalizePackageName(s string) string {
	result := trimNonLetterPrefix(s)
	return strings.ToLower(result)
}
