package string

import "strconv"

func rotatedDigits(N int) int {

	res := 0

	for i := 1; i <= N; i ++ {
		if judegeRotetedNumber(i) {
			res ++
		}
	}
	return res
}

func judegeRotetedNumber(I int) bool {
	By := []byte( strconv.Itoa(I))
	Good := false
	noInvalid := true

	for _,b := range By {
		if b == '2' || b == '5' || b == '6' || b == '9' {
			Good = true
		}
		if b== '3' || b == '4' || b == '7' {
			noInvalid = false
		}
	}
	return Good && noInvalid
}
