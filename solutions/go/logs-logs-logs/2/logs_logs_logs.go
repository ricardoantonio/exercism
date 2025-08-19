package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	const (
		RecommendationIcon = '❗'
		SearchIcon         = 128269 // 🔍
		WeatherIcon        = 9728   // ☀
	)
	for _, char := range log {
		switch char {
		case RecommendationIcon:
			return "recommendation"
		case SearchIcon:
			return "search"
		case WeatherIcon:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var builder strings.Builder
	for _, char := range log {
		if char == oldRune {
			builder.WriteRune(newRune)
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
