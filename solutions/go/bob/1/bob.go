package bob

import "strings"

func Hey(remark string) string {
	endsWithQuestionMark := strings.HasSuffix(strings.TrimSpace(remark), "?")
	allCaps := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	allWhitespace := strings.TrimSpace(remark) == ""

	question := endsWithQuestionMark
	yelling := allCaps
	yellingQuestion := allCaps && endsWithQuestionMark
	silence := allWhitespace

	switch {
	case silence:
		return "Fine. Be that way!"
	case yellingQuestion:
		return "Calm down, I know what I'm doing!"
	case yelling:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	default:
		return "Whatever."
	}
}
