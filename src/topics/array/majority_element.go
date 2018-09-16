package array

// Moore voting algorithm
// it is used when a majority element certainly exist
func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v != candidate {
			count --
		} else {
			count ++
		}
	}
	return candidate
}
