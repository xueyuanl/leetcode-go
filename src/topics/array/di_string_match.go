package array

func diStringMatch(S string) []int {

	I := 0
	D := len(S)

	var res  = make([]int,0, len(S) + 1)
	for _, v := range S {
		if v == 'I' {
			res = append(res, I)
			I ++
		} else {
			res = append(res, D)
			D --
		}
	}
	res = append(res, I)
	return res
}
