package string

func detectCapitalUse(word string) bool {

	OnlyFirst := false

	bytes := []byte(word)

	len := len(bytes)

	if bytes[0] >= 65 && bytes[0] <= 90 {
		OnlyFirst = true
	}

	capNumber := 0
	for i := 1; i < len; i ++ {
		if bytes[i] >= 65 && bytes[i] <= 90 {
			capNumber ++
		}
	}

	if (OnlyFirst && capNumber == len-1) || capNumber == 0 {
		return true
	}

	return false

}
