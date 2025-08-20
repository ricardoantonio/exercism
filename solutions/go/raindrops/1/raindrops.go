package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	raindropSound := ""
	if number%3 == 0 {
		raindropSound += "Pling"
	}
	if number%5 == 0 {
		raindropSound += "Plang"
	}
	if number%7 == 0 {
		raindropSound += "Plong"
	}
	if raindropSound == "" {
		raindropSound = strconv.Itoa(number)
	}
	return raindropSound
}
