package array

func flipAndInvertImage(A [][]int) [][]int {

	for _, row := range A {
		revertRow(row)

	}
	return A
}

func revertRow(row []int) {
	head := 0
	tail := len(row) - 1

	for head <= tail {
		if row[head] == row[tail] {
			row[head] ^= 1
			if head != tail {
				row[tail] ^= 1
			}

		}
		head ++
		tail --
	}
}
