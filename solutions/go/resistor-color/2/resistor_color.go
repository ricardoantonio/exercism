package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
	}
}

var colorMap = func() map[string]int {
	m := make(map[string]int)
	for i, c := range Colors() {
		m[c] = i
	}
	return m
}()

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	if val, ok := colorMap[color]; ok {
		return val
	}
	return -1
}
