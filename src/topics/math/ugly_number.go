package math


func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}

	if num % 2 != 0 || num % 3 != 0 || num % 5 != 0 {
		return false
	}

	return (num % 2 == 0 && isUgly(num / 2)) || (num % 3 == 0 && isUgly(num / 3)) || ( num % 5 == 0 && isUgly( num/ 5))
}