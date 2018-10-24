package bits

func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i ++ {
		if isBinaryPrime(i) {
			res ++
		}
	}
	return res
}

func isBinaryPrime(i int) bool {

	bitsCount := 0

	for i > 0 {
		bitsCount ++
		i = i & (i - 1)    //del the last bit  1 once time
	}
	return isPrime(bitsCount)
}


func isPrime(data int) bool {
	if data == 1 {
		return false
	}
	if data == 2 {
		return true
	}
	if data % 2 == 0 {
		return false
	}
	for i := 3; i * i <= data; i = i + 2 {
		if data % i == 0 {
			return false
		}
	}
	return true
}
