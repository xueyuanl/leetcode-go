package dynamic

type NumArray struct {
	nums []int
	dp [] int
}


func Constructor(nums []int) NumArray {
	var numArray NumArray
	numArray.nums = nums
	numArray.dp = make([]int, len(nums))
	base := 0
	for i, v := range nums {
		numArray.dp[i] = base + v
		base = numArray.dp[i]
	}
	return  numArray
}


func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j] - this.dp[i] + this.nums[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */