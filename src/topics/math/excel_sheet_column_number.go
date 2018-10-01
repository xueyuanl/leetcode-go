package math


func titleToNumber(s string) int {
	len := len(s)
	res := 0
	for i := 0; i < len; i ++ {
		res = res * 26 + (int(s[i]) - 64)
	}
	return res
}
