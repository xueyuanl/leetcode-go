package main

import "fmt"

func main(){


	s := "abcdefg"

	ss := reverseStr(s,2)

	fmt.Println(ss)


}


func reverseStr(s string, k int) string {

	k *= 2
	sli := []byte(s)

	head, tail := 0, len(sli) - 1

	for  head + k - 1 < tail {

		reverse(&sli, head, head + k / 2 - 1)
		head += k

	}

	if head + k / 2 - 1 < tail{
		reverse(&sli, head, head + k / 2 - 1)
	}else {
		reverse(&sli, head, tail)
	}


	return string(sli)
}



func reverse(sli *[]byte, head int, tail int){
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