package bits

func hasAlternatingBits(n int) bool {

	mark := 0

	for n > 0 {

		mark = n & 1
		n = n >> 1

		if n&1 == mark {
			return false
		}
	}

	return true
}
