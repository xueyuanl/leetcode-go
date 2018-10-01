package bits

func singleNumber(nums []int) int {
	var r int
	for _, num := range nums {
		r ^= num
	}
	return r
}

//silly solution using map
func singleNumbers(nums []int) int {

	out := 0
	mapping := make(map[int] int)

	for _,v := range nums {
		if _, ok := mapping[v]; ok {
			mapping[v] = 2
		} else {
			mapping[v] = 1
		}
	}

	for k,v := range mapping {
		if v == 1{
			out = k
			break
		}
	}
	return out
}