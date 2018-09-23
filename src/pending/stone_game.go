package pending


//Time Limit Exceeded
func stoneGame(piles []int) bool {

	sum := 0

	for _, v := range piles {
		sum += v
	}

	alex := getPoints(piles)
	lee := sum - alex

	return alex > lee
}

func getPoints(piles []int) int {
	len := len(piles)
	if len == 0 {
		return 0
	}
	if len == 2 {
		return max(piles[0],piles[1])
	}
	return max(piles[0] + min(getPoints(piles[1:len - 1]), getPoints(piles[2:])),
		 		piles [len - 1] + min(getPoints(piles[:len - 2]), getPoints(piles[1:len - 1])))

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