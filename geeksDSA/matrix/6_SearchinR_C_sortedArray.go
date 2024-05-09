package matrix

func SearchInRowColumnSortedMatrix(mat [][]int, targer int) (int, int) {

	x, y := -1, -1

	for i := len(mat[0]) - 1; i >= 0; i-- {

		if mat[0][i] == targer {
			x = 0
			y = i
			break
		}
		if mat[0][i] < targer {
			y = i
		}

		if y != -1 {
			for j := 1; j < len(mat); j++ {

				if mat[j][y] == targer {
					x = j
					break
				}

			}
		}

	}

	if x == -1 || y == -1 {
		return -1, -1
	}

	return x, y
}
