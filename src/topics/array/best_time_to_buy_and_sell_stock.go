package array

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var buyPoint int = prices[0]

	var maxProfit int

	for _, v := range prices {

		if v > buyPoint && v-buyPoint > maxProfit {
			maxProfit = v - buyPoint
		}

		if v < buyPoint {
			buyPoint = v
		}
	}

	return maxProfit
}
