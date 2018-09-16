package array

func sortArrayByParity(A []int) []int {

	var res []int

	var odd []int

	for _, v := range A {
		if v & 1 == 1 {
			odd = append(odd,v)
		} else {
			res = append(res, v)
		}
	}
	res = append(res,odd...)

	return res
}