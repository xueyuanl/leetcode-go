package dynamic


func maxProfit(prices []int, fee int) int {

	len := len(prices)
	buy := make([]int, len)
	sell := make([]int, len)
	buy[0] = -prices[0]
	for i:= 1; i < len; i ++ {
		buy[i] = max(buy[i - 1], sell[i - 1] - prices[i])
		sell[i] = max(sell[i - 1], buy[i - 1] + prices[i] - fee)
	}

	return sell[len - 1]

}

/*At day i, we may buy stock (from previous sell status) or do nothing (from previous buy status):
buy[i] = Math.max(buy[i - 1], sell[i - 1] - prices[i]);
Or
At day i, we may sell stock (from previous buy status) or keep holding (from previous sell status):
sell[i] = Math.max(sell[i - 1], buy[i - 1] + prices[i]);*/