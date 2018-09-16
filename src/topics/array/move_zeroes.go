package array

func moveZeroes(nums []int) {

	len := len(nums)

	copy := make([]int, len)

	var index int
	for i := 0; i < len; i ++ {
		if nums[i] != 0 {
			copy[index ] = nums[i]
			index ++
		}
	}

	for i := 0; i < len; i ++ {
		nums[i] = copy[i]
	}
}
