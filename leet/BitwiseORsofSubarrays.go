package main

import "fmt"

func main() {

	sli := []int{1, 1, 2}

	subarrayBitwiseORs(sli)

}

func subarrayBitwiseORs(A []int) int {

	if len(A) == 1 {
		return 1
	}

	sli := [] int{}
	for i := 0; i < len(A); i ++ {

		var or = 0

		for j := i; j < len(A); j ++ {
			or = or | A[j]
			sli = append(sli, or)
		}

	}

	fmt.Println(&sli, sli)


	UniqueSlice(&sli)

	return len(sli)
}

func UniqueSlice(slice *[]int) {
	found := make(map[int]bool)
	total := 0
	for i, val := range *slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(*slice)[total] = (*slice)[i]
			total++
		}
	}

	*slice = (*slice)[:total]
}

func RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		return RemoveRepByLoop(slc)
	} else {
		return RemoveRepByMap(slc)
	}
}

func RemoveRepByLoop(slc []int) []int {
	result := []int{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

func RemoveRepByMap(slc []int) []int {
	var result []int

	tempMap := map[int]byte{}

	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
