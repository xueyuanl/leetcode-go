package src

import "./stringutil"

func main() {

}

func reverseWords(s string) string {

	sli := []byte(s)

	space := byte(' ')

	head, index, len := 0, 0, len(sli)

	for ; index < len && head < len; index ++ {

		if sli[index] == space {
			stringutil.Reverse(&sli, head, index-1)
			head = index + 1
		} else if index == len-1 {
			stringutil.Reverse(&sli, head, index)
		}

	}

	return string(sli)
}
