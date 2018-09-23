package dynamic

func isSubsequence(s string, t string) bool {

	ss := []byte(s)
	tt := []byte(t)
	if len(ss) == 0 {
		return true
	}
	index := 0
	for _, v := range tt {
		if v == ss[index]{
			index ++
		}
	}

	return index == len(ss)
}
