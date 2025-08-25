package isbn

func IsValidISBN(isbn string) bool {
	sum := 0
	position := 0

	for _, char := range isbn {
		if char == '-' {
			continue
		}

		var value int
		if char >= '0' && char <= '9' {
			value = int(char - '0')
		} else if (char == 'X' || char == 'x') && position == 9 {
			value = 10
		} else {
			return false
		}

		position++
		if position > 10 {
			return false
		}

		sum += value * (10 - (position - 1))
	}

	if position != 10 {
		return false
	}

	return sum%11 == 0
}
