package hash

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	As := strings.Split(A, " ")
	Bs := strings.Split(B, " ")

	var res []string

	hash := make(map[string]int)

	for _, v := range As {
		if _, ok := hash[v]; ok {
			hash[v] = hash[v] + 1
		} else {
			hash[v] = 1
		}
	}

	for _, k := range Bs {
		if _, ok := hash[k]; ok {
			hash[k] = hash[k] + 1
		} else {
			hash[k] = 1
		}
	}

	for k, v := range hash {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}
