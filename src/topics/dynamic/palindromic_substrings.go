package dynamic

//#classic

func countSubstrings(s string) int {
	var b []byte = []byte(s)
	count := 0

	for i, _ := range b {
		count += getPalindromic(b, i, i)
		count += getPalindromic(b,i,i + 1)
	}
	return count
}

func getPalindromic(b []byte, i int, j int) int {
	count := 0
	for i >= 0 && j < len(b) && b[i] == b[j] {
		count ++
		i --
		j ++
	}
	return count
}
