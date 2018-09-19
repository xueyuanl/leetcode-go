package dynamic

//dp[i] = cost[i] + min(dp[i - 1], dp[i - 2])
func minCostClimbingStairs(cost []int) int {
	len := len(cost)
	if len == 2 {
		return min(cost[0], cost[1])
	}
	var dp [] int = make([]int, len)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i:= 2; i < len ;i ++ {
		dp[i] = cost[i] + min(dp[i - 1], dp[i - 2])
	}

	return min(dp[len - 1], dp[len - 2])
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}