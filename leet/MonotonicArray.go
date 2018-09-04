package main

func main() {
	
}


func isMonotonic(A []int) bool {

	result1 := true
	result2 := true

	for i := 0; i < len(A); i ++ {
		if i + 1 != len(A) && A[i] > A[i +1]{
			result1 = false
		}
	}

	for i := 0; i < len(A); i ++ {
		if i + 1 != len(A) && A[i] < A[i +1]{
			result2 = false
		}
	}

	if result1 == true || result2 == true{
		return true
	}else {
		return false
	}

}