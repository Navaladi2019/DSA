package matrix

func FindSpiralPattern(mat [][]int) []int {

	arr := make([]int, 0, len(mat)*len(mat[0]))
	left := 0                // column
	right := len(mat[0]) - 1 // Column
	top := 0                 // Row
	bottom := len(mat) - 1   // Row

	for top <= bottom && left <= right {

		// appending top
		for l := left; l <= right; l++ {
			arr = append(arr, mat[top][l])
		}

		top++

		// appending right
		for t := top; t <= bottom; t++ {

			arr = append(arr, mat[t][right])
		}

		right--

		if top <= bottom {
			// appending bottom
			for r := right; r >= left; r-- {
				arr = append(arr, mat[bottom][r])
			}

			bottom--
		}

		if left <= right {
			// appending left
			for b := bottom; b >= top; b-- {

				arr = append(arr, mat[b][left])
			}

			left++
		}

	}

	return arr
}
