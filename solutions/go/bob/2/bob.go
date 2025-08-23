package bob

import "strings"

func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)

	if trimmed == "" {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(trimmed, "?")
	isYelling := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark

	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
