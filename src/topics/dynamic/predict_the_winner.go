package dynamic

func PredictTheWinner(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	player1 := getScores(nums)
	player2 := sum - player1

	return player1 >= player2

}

func getScores(nums []int) int {
	len := len(nums)

	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		return max(nums[0], nums[1])
	}

	return max(nums[0]+min(getScores(nums[2:len]), getScores(nums[1:len-1])),
		nums[len-1]+min(getScores(nums[:len - 2]), getScores(nums[1:len - 1])))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
