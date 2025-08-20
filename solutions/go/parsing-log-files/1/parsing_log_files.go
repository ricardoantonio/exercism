package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	totalQuotedPasswords := 0
	for _, line := range lines {
		if re.MatchString(line) {
			totalQuotedPasswords++
		}
	}
	return totalQuotedPasswords
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	for index, line := range lines {
		sl := re.FindStringSubmatch(line)
		if sl == nil {
			continue
		}
		lines[index] = "[USR] " + sl[1] + " " + line
	}
	return lines
}
