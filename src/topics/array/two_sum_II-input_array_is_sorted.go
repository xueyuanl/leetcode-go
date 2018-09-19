package array

func twoSum(numbers []int, target int) []int {
	len := len(numbers)
	l, r := 0, len-1
	for l < r {
		for numbers[l]+numbers[r] > target {
			r --
		}
		for numbers[l]+numbers[r] < target {
			l ++
		}
		if numbers[l]+numbers[r] == target {
			break
		}
	}
	var res []int
	res = append(res, l+1, r+1)
	return res
}

//O(n*n) bad solution
func twoSums(numbers []int, target int) []int {
	var result []int

lable:
	for i, v := range numbers {
		for j, k := range numbers[i+1:] {
			if v+k > target {
				break
			}
			if v+k == target {
				result = append(result, i+1, i+j+2)
				break lable
			}
		}
	}

	return result
}
