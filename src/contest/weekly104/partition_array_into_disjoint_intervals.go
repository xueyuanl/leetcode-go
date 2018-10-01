package weekly104


func partitionDisjoint(A []int) int {
	len := len(A)

	leftMax := make([]int, len)
	rightMin := make([]int, len)

	leftMax[0] = A[0]
	for i := 1 ; i< len ; i ++ {
		if A[i] > leftMax[i - 1] {
			leftMax[i] = A[i]
		} else {
			leftMax[i] = leftMax[i - 1]
		}
	}

	rightMin[len - 1] = 1000000 + 1
	for i := len - 2; i >=0; i -- {
		if A[i + 1] < rightMin[i + 1] {
			rightMin[i] = A[i + 1]
		}else {
			rightMin[i] = rightMin[i + 1]
		}
	}

	res := 0
	for i := 0 ; i< len ; i ++ {
		if leftMax[i] <= rightMin[i]  {
			res = i + 1
			break
		}
	}
	return res
}