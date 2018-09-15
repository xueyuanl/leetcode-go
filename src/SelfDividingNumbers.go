package src

import "fmt"

func main() {

	left, right := 101, 102

	result := selfDividingNumbers(left, right)

	fmt.Println(result)
}

func selfDividingNumbers(left int, right int) []int {

	var result []int

	for i := left; i <= right; i ++ {

		if isSelfDividing(i) {
			result = append(result, i)
		}

	}

	return result
}

func isSelfDividing(v int) bool {

	if(v > 9 && v % 10 == 0){
		 return  false
	}

	y := v

	for y%10 != 0 {

		p := y % 10

		if v%p != 0 {
			return false
		}

		y = y / 10

		if y != 0 && y % 10 == 0{
			return false
		}
	}

	return true

}
