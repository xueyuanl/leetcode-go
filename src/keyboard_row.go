package src

import (
	"regexp"
	"strings"
)

func findWords(words []string) []string {

	var out []string
	for _, v := range words {
		if inOneRow(v) {
			out = append(out, v)
		}
	}
	return out
}

func inOneRow(s string) bool {
	lower := strings.ToLower(s)
	pattern := "^[qwertyuiop]*$|^[asdfghjkl]*$|^[zxcvbnm]*$"
	rex := regexp.MustCompile(pattern)
	return rex.MatchString(lower)
}
