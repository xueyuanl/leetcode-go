package main

func main() {



	isLongPressedName("znxnorutzt", "zznxxnnooruuzztt")
}


func isLongPressedName(name string, typed string) bool {

	nameby := []byte(name)
	typedby := []byte(typed)

	if nameby[0] != typedby[0] {
		return false
	}

	tpyedIndex := 0

	for i := 1 ; i < len(nameby); i ++ {
		tpyedIndex ++

		if typedby[tpyedIndex] != nameby[i] && typedby[tpyedIndex] != typedby[tpyedIndex - 1] && tpyedIndex > len(typedby) - 1{
			return false
		}

		for typedby[tpyedIndex] != nameby[i] && typedby[tpyedIndex] == typedby[tpyedIndex - 1] {
			tpyedIndex ++
			if tpyedIndex > len(typedby) - 1 {
				return false
			}
		}
		if typedby[tpyedIndex] != nameby[i] {
			return false
		}
	}
	return true
}