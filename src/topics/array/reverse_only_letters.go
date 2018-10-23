package array

func reverseOnlyLetters(S string) string {

	str := []byte(S)

	begin, end := 0, len(str) - 1

	for begin < end {
		for !isAlph(str[begin]) && begin < end {
			begin ++
		}
		for !isAlph(str[end]) && begin < end {
			end --
		}
		if begin < end {
			str[begin] = str[begin] ^ str[end]
			str[end] = str[begin] ^ str[end]
			str[begin] = str[begin] ^ str[end]
			begin ++
			end --
		}
	}

	return string(str)
}

func isAlph(b byte) bool {
	return b > 64 && b < 91 || b > 96 && b < 123
}