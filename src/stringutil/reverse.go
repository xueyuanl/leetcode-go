package stringutil


func Reverse(sli *[]byte, head int, tail int){
	if head >= tail {
		return
	}
	for head < tail {

		exchange(sli, head, tail)

		head ++
		tail --
	}
}

func exchange(sli *[]byte, head int, tail int) {
	temp := (*sli)[head]
	(*sli)[head] = (*sli)[tail]
	(*sli)[tail] = temp
}