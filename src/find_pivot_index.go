package src



func pivotIndex(nums []int) int {

	var leftSum, rightSum int

	for _,v := range nums {
		rightSum += v
	}

	for i,v := range nums {
		if i != 0 {
			leftSum += nums[i - 1]
		}
		rightSum -= v

		if leftSum == rightSum {
			return i
		}

	}

	return -1
}