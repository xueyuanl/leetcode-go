package src


func findDuplicates(nums []int) []int {

	mapping := make([]int, len(nums) + 1)

	var twice [] int
	for _, v := range nums {

		if mapping[v] != 0 {
			twice = append(twice, v)
		} else {
			mapping[v] ++
		}

	}
	return twice

}