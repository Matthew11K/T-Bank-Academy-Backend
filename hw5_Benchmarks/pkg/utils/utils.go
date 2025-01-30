package utils

import (
	"regexp"
	"strings"
)

func CountSubstringsStringsCount(s, substr string) int {
	return strings.Count(s, substr)
}

func CountSubstringsRegex(s, substr string) (int, error) {
	pattern := regexp.QuoteMeta(substr)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return 0, err
	}

	return len(re.FindAllStringIndex(s, -1)), nil
}

func CountSubstringsManual(s, substr string) int {
	count := 0
	index := 0

	for {
		idx := strings.Index(s[index:], substr)
		if idx == -1 {
			break
		}

		count++
		index += idx + len(substr)
	}

	return count
}
