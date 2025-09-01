package romannumerals

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	symbol string
	value  int
}

var romanNumerals = []romanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

var ErrOutOfRange = errors.New("input must be between 1 and 3999")

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", ErrOutOfRange
	}

	var sb strings.Builder
	for _, numeral := range romanNumerals {
		count := input / numeral.value
		if count > 0 {
			sb.WriteString(strings.Repeat(numeral.symbol, count))
			input = input % numeral.value
		}
	}
	return sb.String(), nil
}
