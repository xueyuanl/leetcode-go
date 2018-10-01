package string

import "strings"

func findLUSlength(a string, b string) int {

	strings.Compare()
	if a == b {
		return -1
	}

	lena, lenb := len(a), len(b)

	if lena < lenb {
		return lenb
	}
	return lena
}