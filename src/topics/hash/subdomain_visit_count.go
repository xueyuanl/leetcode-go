package hash

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	visited := make(map[string]int)

	for _,v := range cpdomains {
		var numAndDomains []string
		numAndDomains = strings.Split(v, " ")

		num,_ := strconv.Atoi(numAndDomains[0])
		domains := numAndDomains[1]

		for len(domains) > 0 {
			if _,ok := visited[domains]; ok {
				visited[domains] += num
			}else {
				visited[domains] = num
			}
			domains = cutSubDomain(domains)
		}
	}

	var res []string
	for k,v := range visited {
		s := strconv.Itoa(v) + " " + k
		res = append(res, s)
	}
	return res
}

func cutSubDomain(s string) string {
	var res []byte
	slices := []byte(s)
	for i,v := range slices {
		if v == '.' {
			res = slices[i + 1:]
			return string(res)
		}
	}
	return ""
}