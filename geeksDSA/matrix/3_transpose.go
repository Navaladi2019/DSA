package matrix

func FindTranspose(mat [][]int) {

	for i := 0; i < len(mat); i++ {
		for j := i + 1; j < len(mat[i]); j++ {
			mat[j][i], mat[i][j] = mat[i][j], mat[j][i]
		}
	}
}

func FindTranspose_1(mat [][]int) {

	for i := 0; i < len(mat); i++ {
		for j := i; j < len(mat[i]); j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
}
