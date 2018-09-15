package src

import "fmt"

func main() {

	A := []int{1, 2, 4}

	result := peakIndexInMountainArray(A)

	fmt.Println(result)

}

func peakIndexInMountainArray(A []int) int {

	len := len(A)

	if len == 1 {
		return 0
	}
	for i, v := range A {

		if i+1 < len && v > A[i+1] {
			return i
		}
	}

	return len - 1
}
