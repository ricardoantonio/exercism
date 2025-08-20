package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",  // 0
		"brown",  // 1
		"red",    // 2
		"orange", // 3
		"yellow", // 4
		"green",  // 5
		"blue",   // 6
		"violet", // 7
		"grey",   // 8
		"white",  // 9
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, c := range Colors() {
		if c == color {
			return i
		}
	}
	return -1
}
