package isogram

import "unicode"

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)

	for _, char := range word {
		char = unicode.ToLower(char)
		if !unicode.IsLetter(char) {
			continue
		}
		if letters[char] {
			return false
		}
		letters[char] = true
	}
	return true
}
