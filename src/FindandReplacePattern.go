package src



func findAndReplacePattern(words []string, pattern string) []string {
	var out [] string

	mappedPattern := mapStr(pattern)

	for _,v := range words {

		mappedV := mapStr(v)

		if stringSliceEqual(mappedPattern, mappedV) {
			out = append(out, v)
		}
	}

	return out
}

func mapStr( strs string) []int{
	var str [] byte = []byte(strs)
	k := 1

	m := make(map[byte] int)

	for _, v := range str {
		_, exist := m[v]
		if ! exist {
			m[v] = k
			k ++
		}
	}

	var out [] int

	for _,v := range str {
		out = append(out, m[v])
	}

	return out

}

func stringSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}