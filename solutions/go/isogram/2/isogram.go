package isogram

func IsIsogram(s string) bool {
	var seen uint32
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c |= 32
		}
		if 'a' <= c && c <= 'z' {
			bit := uint32(1) << (c - 'a')
			if seen&bit != 0 {
				return false
			}
			seen |= bit
		}
	}
	return true
}
