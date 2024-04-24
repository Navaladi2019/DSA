package matrix

func Rotateby90Naive(mat [][]int) [][]int {

	opMaztrix := make([][]int, len(mat))

	for i := 0; i < len(opMaztrix); i++ {
		opMaztrix[i] = make([]int, len(mat))
	}

	for i := 0; i < len(mat); i++ {

		colEle := len(mat) - 1 - i
		for j := 0; j < len(mat); j++ {

			opMaztrix[i][j] = mat[j][colEle]
		}

	}
	return opMaztrix
}

func RotateBy90(mat [][]int) {

	// transpose the matrix and reverse the individual column

	FindTranspose(mat)

	for i := 0; i < len(mat); i++ {

		low := 0
		high := len(mat) - 1

		for low != high {

			mat[low][i], mat[high][i] = mat[high][i], mat[low][i]

			low++
			high--
		}
	}
}
