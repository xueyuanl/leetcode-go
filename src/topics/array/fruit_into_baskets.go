package array

//bullshit solution
func totalFruit(tree []int) int {

	type1, type2 := -1, -1

	len := len(tree)


	res, k := 0, 0

	newBegin := -1

	for i:= 0; i < len; i ++ {

		if type1 == -1 {
			type1 = tree[i]
		} else if type1 != -1 && type2 == -1 && tree[i] != type1 {
			type2 = tree[i]
			newBegin = i
		}

		k ++

		if i + 1 < len && type1 != -1 && type2 != -1 && tree[i + 1] != type1 && tree[i + 1] != type2 {

			if k > res {
				res = k
			}
			k = 1
			type1 = tree[i]
			type2 = -1

			if newBegin != -1 {
				k = 1
				type1 = tree[newBegin]
				type2 = -1
				i = newBegin
			}
			newBegin = i  + 1
		}

		if i == len - 1  && k > res{
			res = k
		}


	}
	return res

}