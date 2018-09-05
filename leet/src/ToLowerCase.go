package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	s = "HelloO"

	k := toLowerCase(s)

	fmt.Println(k)

}


func toLowerCase(str string) string {

	sli := []byte(str)


	s := strings.ToLower(str)

	return s
}

/*A better solution*/
/*func toLowerCase(str string) string {
	r:=[]rune(str)
	for i:=0;i<len(r);i++{
		if r[i]<=90&&r[i]>=65 {
			r[i]+=32
		}
	}
	return string(r)
}*/