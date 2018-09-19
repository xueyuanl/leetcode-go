package dynamic

//using now record the max!!!
func rob(nums []int) int {
	last, now, temp:= 0, 0, 0

	for _, v := range nums {
		temp = now
		now = max(last + v, now)
		last = temp
	}

	return now
}


//record the history status
func robs(nums []int) int {
	len := len(nums)
	if len == 0 {
		return  0
	}
	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return max(nums[0], nums[1])
	}

	maxs := 0

	dp := make([]int, len)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len ;i ++ {
		dp[i] = max(dp[i - 1], nums[i] + dp[i - 2])
		if dp[i] > maxs {
			maxs = dp[i]
		}
	}
	return maxs
}

func max(a int, b int) int {
	if a > b {
		 return a
	}
	return b
}