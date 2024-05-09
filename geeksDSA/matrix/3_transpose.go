package matrix

func FindTranspose(mat [][]int) {

	for i := 0; i < len(mat); i++ {
		for j := i + 1; j < len(mat[i]); j++ {
			mat[j][i], mat[i][j] = mat[i][j], mat[j][i]
		}
	}
}
