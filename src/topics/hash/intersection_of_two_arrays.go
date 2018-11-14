package hash


func intersection(nums1 []int, nums2 []int) []int {

	deDupMap := make(map[int] bool)

	for _,v := range nums1 {
		if ! deDupMap[v] {
			deDupMap[v] = true
		}
	}

	resMap := make(map[int] bool)

	for _,v := range nums2 {

		if deDupMap[v] {
			resMap[v] = true
		}

	}

	var res []int

	for k,_ := range resMap {

		res = append(res, k)

	}
	return res

}