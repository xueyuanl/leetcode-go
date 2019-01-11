package array

func sortArrayByParityII(A []int) []int {

	odd, even := 0, 0
	var res []int
	current := 0

	for current < len(A) {

		if current&1 == 1 { // odd
			for A[odd]&1 != 1 {
				odd ++
			}
			res = append(res, A[odd])
			odd ++
		} else {

			for A[even]&1 != 0 {
				even ++
			}
			res = append(res, A[even])
			even ++
		}
		current ++
	}
	return res
}
