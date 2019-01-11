package string

var morseStrs = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
var morseMap = map[byte]string{}

func uniqueMorseRepresentations(words []string) int {

	for i, v := range morseStrs {
		morseMap[byte('a'+i)] = v
	}

	resMap := make(map[string]bool)
	res := 0
	for _, v := range words {

		morseStr := getMorseStr(v)
		if !resMap[morseStr] {
			resMap[morseStr] = true
			res ++
		}
	}
	return res
}

func getMorseStr(word string) string {

	res := ""
	for _,v := range word {

		res += morseMap[byte(v)]
	}
	return res
}
