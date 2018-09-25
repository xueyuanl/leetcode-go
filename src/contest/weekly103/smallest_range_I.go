package weekly103

func smallestRangeI(A []int, K int) int {

	min, max := 10000, 0
	for _, v := range A {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	diff := max - min

	if diff-2*K < 0 {
		return 0
	} else {
		return diff - 2*K
	}

}
