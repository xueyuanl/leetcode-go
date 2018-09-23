package dynamic

//classic dp problem
func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	dp := make([]int, len(A))
	sum := 0
	for i := 2; i < len(A); i ++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i - 1] + 1
			sum += dp[i]
		}
	}

	return sum
}
