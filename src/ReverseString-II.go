package src

import "fmt"
import "./stringutil"

func main() {

	s := "abcdefg"

	ss := reverseStr(s, 2)

	fmt.Println(ss)

}

func reverseStr(s string, k int) string {

	k *= 2
	sli := []byte(s)

	head, tail := 0, len(sli)-1

	for head+k-1 < tail {

		stringutil.Reverse(&sli, head, head+k/2-1)
		head += k

	}

	if head+k/2-1 < tail {
		stringutil.Reverse(&sli, head, head+k/2-1)
	} else {
		stringutil.Reverse(&sli, head, tail)
	}

	return string(sli)
}
