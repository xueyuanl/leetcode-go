package src

func maxIncreaseKeepingSkyline(grid [][]int) int {

	column_len := len(grid)
	row_len := len(grid[0])

	column_array := [51] int{}
	row_array := [51] int{}

	for i := 0; i < column_len; i ++ {
		for j := 0; j < row_len; j ++ {

			if grid[i][j] > row_array[j] {
				row_array[j] = grid[i][j]
			}
			if grid[i][j] > column_array[i] {
				column_array[i] = grid[i][j]
			}
		}
	}

	sum := 0
	for i := 0; i < column_len; i ++ {
		for j := 0; j < row_len; j ++ {
			if column_array[i] > row_array[j]{
				sum += row_array[j] - grid[i][j]
			} else{

				sum += column_array[i] - grid[i][j]
			}

		}

	}
	return  sum

}
