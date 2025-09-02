package matrix

func GetArrayInSnakePattern(mat [][]int) (arr []int) {

	for i := 0; i < len(mat); i++ {

		if i%2 == 0 {
			arr = append(arr, mat[i]...)

		} else {

			for j := len(mat[i]) - 1; j >= 0; j-- {
				arr = append(arr, mat[i][j])
			}

		}

	}

	return arr
}

func GetArrayInSnakePattern_1(mat [][]int) (arr []int) {

	isLeft := true

	res := make([]int, 0, len(mat)*len(mat[0]))

	for i := 0; i < len(mat); i++ {
		if isLeft {
			res = append(res, mat[i]...)
			isLeft = false
		} else {
			for j := len(mat[i]) - 1; j >= 0; j-- {
				res = append(res, mat[i][j])
			}
			isLeft = true
		}
	}

	return res
}
