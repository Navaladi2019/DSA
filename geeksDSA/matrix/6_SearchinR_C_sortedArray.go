package matrix

func SearchInRowColumnSortedMatrix(mat [][]int, target int) (int, int) {

	x, y := -1, -1

	for i := len(mat[0]) - 1; i >= 0; i-- {

		if mat[0][i] == target {
			x = 0
			y = i
			break
		}
		if mat[0][i] < target {
			y = i
		}

		if y != -1 {
			for j := 1; j < len(mat); j++ {

				if mat[j][y] == target {
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

func SearchInRowColumnSortedMatrix_1(mat [][]int, target int) (x int, y int) {

	x = -1
	y = -1

	for i := 0; i < len(mat[0]); i++ {

		// target found in first row itself
		if mat[0][i] == target {
			x = 0
			y = i
			return
		}

		if mat[0][i] < target {

			for j := 0; j < len(mat); j++ {

				if mat[j][i] == target {

					x = j
					y = i
					return
				} else if mat[j][i] > target {
					break
				}
			}
		} else {
			break
		}

	}

	return

}

func SearchInRowColumnSortedMatrix_2(mat [][]int, target int) (x int, y int) {

	x = -1
	y = -1

	defer func() {
		if x == -1 || y == -1 {
			x = -1
			y = -1
		}

	}()

	for i := 0; i < len(mat[0]); i++ {

		if mat[0][i] == target {

			x = 0
			y = i
			return
		} else if mat[0][i] > target {

			x = -1
			y = -1
			break
		}

		if mat[0][i] > target {
			break
		}

		y = i

		for j := 1; j < len(mat); j++ {

			if mat[j][y] == target {
				x = j
				return
			}

			if mat[j][y] > target {
				break
			}
		}

	}

	return
}
