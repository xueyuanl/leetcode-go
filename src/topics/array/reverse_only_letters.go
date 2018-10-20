package main


func main () {

	reverseOnlyLetters("(;,9=/'*")
}


func reverseOnlyLetters(S string) string {

	str := []rune(S)

	begin, end := 0, len(str) - 1

	for begin < end {
		for !isAlph(str[begin]) && begin < end {
			begin ++
		}
		for !isAlph(str[end]) && begin < end {
			end --
		}
		str[begin] = str[begin] ^ str[end]
		str[end] = str[begin] ^ str[end]
		str[begin] = str[begin] ^ str[end]
		begin ++
		end --
	}

	return string(str)
}

func isAlph(b rune) bool {
	return b > 64 && b < 91 || b > 96 && b < 123
}