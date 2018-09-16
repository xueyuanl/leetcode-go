package pending


func findComplement(num int) int {


	var x uint32  = uint32(num)

	return int(^x)



}