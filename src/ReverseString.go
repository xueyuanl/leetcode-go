package src

import "fmt"

func main() {


	s:= "abc, AA@@#$"

	fmt.Println(reverseString(s))

}




func reverseString(s string) string {

	var result []byte

	for i := len(s) - 1; i >= 0; i --{
		result = append(result,s[i])
	}

	return string(result)
}
