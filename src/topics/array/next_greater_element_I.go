package array

//silly solution
func nextGreaterElement(findNums []int, nums []int) []int {

	var out []int
	len := len(nums)
	for _,v := range findNums {
		for j,k := range nums {
			if v == k {
				out = append(out, getGreaterElement(v,nums[j:len] ))
			}
		}
	}
	return out
}

func getGreaterElement( v int, nums [] int)  int{
	for _,k := range nums {
		if k > v {
			return k
		}
	}
	return -1
}