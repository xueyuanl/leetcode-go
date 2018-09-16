package array


func findMaxConsecutiveOnes(nums []int) int {
	var sum, max int
	for _,v := range nums {

		sum = sum + v
		sum = sum * v

		if sum > max {
			max = sum
		}
	}
	return max
}