package pangram

func IsPangram(input string) bool {
	allLettersMask := (1 << 26) - 1
	lettersFoundMask := 0

	for _, char := range input {
		if 'A' <= char && char <= 'Z' {
			char += ('a' - 'A')
		}
		if 'a' <= char && char <= 'z' {
			index := char - 'a'
			lettersFoundMask |= 1 << index
			if lettersFoundMask == allLettersMask {
				return true
			}
		}

	}
	return lettersFoundMask == allLettersMask
}
