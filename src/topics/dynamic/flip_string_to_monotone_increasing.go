package main

func main() {

	minFlipsMonoIncr("00110")

}

func minFlipsMonoIncr(S string) int {

	one, zero := 0, 0

	for _, v := range S {

		if v == '0' {
			one = 1 + min(one, zero)
			zero = zero
		} else {
			one = min(one, zero) // this two lines 
			zero = 1 + zero

		}
	}

	return min(one, zero)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
