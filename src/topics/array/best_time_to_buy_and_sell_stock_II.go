package array


func maxProfits(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	buyPoint := prices[0]
	profit := 0

	for _, v := range prices[1:] {
		if v < buyPoint {
			buyPoint = v
		}

		if v > buyPoint {
			profit += v - buyPoint
			buyPoint = v
		}
	}

	return profit
}