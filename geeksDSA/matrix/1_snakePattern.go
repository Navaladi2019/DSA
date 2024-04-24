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
