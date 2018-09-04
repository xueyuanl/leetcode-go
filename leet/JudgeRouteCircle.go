package main

func main() {


}

func judgeCircle(moves string) bool {

	sli := [] rune(moves)

	row, cloumn := 0, 0

	for _, v := range sli {

		switch v {
		case 'U':
			cloumn ++
		case 'D':
			cloumn --
		case 'L':
			row --
		case 'R':
			row ++

		}
	}

	if row == 0 && cloumn == 0 {
		return true
	} else {
		return false
	}
}
