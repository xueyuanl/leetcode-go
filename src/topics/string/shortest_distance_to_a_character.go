package string


func shortestToChar(S string, C byte) []int {

	s := []byte(S)

	len := len(s)

	output := make([]int, len)
	//set the distance to the max first
	for i := 0; i < len; i ++ {
		output[i] = 10001
	}


	for i,v := range s {
		if v == C {
			update(i, s, &output)
		}
	}
	return output
}

func update(p int, s []byte, output *[]int) {
	for i := p; i >= 0 ; i -- {
		if (p - i) < (*output)[i] {
			(*output)[i] = p - i
		}
	}

	for k := p + 1; k < len(s); k ++ {
		if k - p < (*output)[k] {
			(*output)[k] = k - p
		}
	}
}