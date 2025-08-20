package luhn

func Valid(id string) bool {
	sum := 0
	// The number of digits processed to get the odd/even position
	n := 0

	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]

		// Skip spaces
		if c == ' ' {
			continue
		}

		// Not valid if contains non-digit characters
		if c < '0' || c > '9' {
			return false
		}

		s := int(c - '0')
		if n%2 == 1 {
			s *= 2
			if s > 9 {
				s -= 9
			}
		}
		sum += s
		n++
	}
	return n > 1 && sum%10 == 0
}
