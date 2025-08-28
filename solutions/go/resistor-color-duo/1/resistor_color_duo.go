package resistorcolorduo

var Colors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) == 0 {
		return 0
	}

	if len(colors) == 1 {
		return Colors[colors[0]]
	}

	return Colors[colors[0]]*10 + Colors[colors[1]]
}
