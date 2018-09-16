package array


func arrayPairSum(nums []int) int {
	var sum int
	sort(nums)

	for i := 0; i < len(nums); i  = i + 2 {
		sum += nums[i]
	}

	return sum
}

func sort(nums []int) {
	var temp []int = make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, temp)
}

func mergeSort(nums []int, l int, r int, temp []int) {
	if !(l < r) { // critical step
		return
	}
	m := (l + r) / 2

	mergeSort(nums, l, m, temp)
	mergeSort(nums, m+1, r, temp)

	merge(nums, l, m+1, r, temp)
}

func merge(nums []int, l int, m int, r int, temp []int) {
	i := l;
	j := m;
	t := l;
	for i < m && j <= r {
		if nums[i] < nums[j] {
			temp[t] = nums[i]
			t ++
			i ++
		} else {
			temp[t] = nums[j]
			t ++
			j ++
		}
	}
	for i < m {
		temp[t] = nums[i]
		t ++
		i ++
	}

	for j <= r {
		temp[t] = nums[j]
		t ++
		j ++
	}

	//copy to original []int
	for l <= r {
		nums[l] = temp[l]
		l ++
	}

}
