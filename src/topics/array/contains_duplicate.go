package array

func containsDuplicate(nums []int) bool {

	var mapping map[int]bool = make(map[int]bool)

	for _, v := range nums {
		if mapping[v] {
			return true
		}
		mapping[v] = true
	}
	return false
}