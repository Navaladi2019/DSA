package matrix

func BoundaryTraversal(mat [][]int) (arr []int) {

	// appending first row
	arr = append(arr, mat[0]...)

	//appending middle last element

	if len(mat) > 2 {
		for i := 1; i <= len(mat)-2; i++ {
			arr = append(arr, mat[i][len(mat[i])-1])
		}
	}

	if len(mat) > 1 {
		matrixlateElement := len(mat) - 1
		lastElementLength := len(mat[matrixlateElement])
		// appending last row
		for i := lastElementLength - 1; i >= 0; i-- {

			arr = append(arr, mat[matrixlateElement][i])
		}

	}

	if len(mat[0]) > 1 {
		// appending middle first element
		for i := len(mat) - 2; i > 0; i-- {

			arr = append(arr, mat[i][0])
		}
	}

	return arr
}
