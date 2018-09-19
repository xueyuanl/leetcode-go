package array

//recurrence
func isOneBitCharacter(bits []int) bool {
	len := len(bits)
	if len == 1 {
		return true
	}

	return cutTailCheck(bits[:len-1])
}

func cutTailCheck(bits []int) bool {
	len := len(bits)
	if len == 1 {
		if bits[0] == 0 {
			return true
		} else {
			return false
		}
	}
	if len == 2 && ((bits[0] == 1 && bits[1] == 0) || (bits[0] == 1 && bits[1] == 1)) {
		return true
	}

	res1, res2 := false, false
	if bits[len-1] == 0   {
		res1 = cutTailCheck(bits[:len-1])
	}
	if (bits[len - 2] == 1 && bits[len - 1] == 0) || (bits[len - 2] == 1 && bits[len - 1] == 1) {
		res2 = cutTailCheck(bits[:len-2])
	}

	return res1 || res2
}
