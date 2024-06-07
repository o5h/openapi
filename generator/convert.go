package generator

import "unicode"

func toPascalCase(s string) string {
	var result string
	capitalizeNext := true
	for _, char := range s {
		if !isAlphaNumeric(char) {
			capitalizeNext = true
		} else {
			if capitalizeNext {
				result += string(unicode.ToUpper(char))
				capitalizeNext = false
			} else {
				result += string(unicode.ToLower(char))
			}
		}
	}
	return result
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
