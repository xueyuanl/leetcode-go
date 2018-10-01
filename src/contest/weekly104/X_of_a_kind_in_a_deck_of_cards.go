package weekly104

func hasGroupsSizeX(deck []int) bool {
	mapping := make(map[int]int)

	max := 0
	for _, v := range deck {
		if k, ok := mapping[v]; ok {
			mapping[v] = k + 1
		} else {
			mapping[v] = 1
		}
		if mapping[v] > max {
			max = mapping[v]
		}
	}

	res := 0
	for _, value := range mapping {
		res = gcd(value, res)
	}

	return res > 1
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
