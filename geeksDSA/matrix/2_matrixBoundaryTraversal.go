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
		matrixlastElement := len(mat) - 1
		lastElementLength := len(mat[matrixlastElement])
		// appending last row
		for i := lastElementLength - 1; i >= 0; i-- {

			arr = append(arr, mat[matrixlastElement][i])
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

func BoundaryTraversal_1(mat [][]int) (arr []int) {

	res := make([]int, 0, 10)

	res = append(res, mat[0]...)

	// if there is only single row in matrix no need to continue
	if len(mat) == 1 {
		return res
	}

	// append all right most column in all rows
	for i := 1; i < len(mat); i++ {
		res = append(res, mat[i][len(mat[i])-1])
	}

	// append all bottom most column of last row in reverse order
	//  except the last column as its already appened by above code
	for i := len(mat[0]) - 2; i >= 0; i-- {
		res = append(res, mat[len(mat)-1][i])
	}

	// append only the left most column i all row except first and last row
	// only if the colmn length is more than one
	if len(mat[0]) > 1 {
		for i := len(mat) - 2; i > 0; i-- {
			res = append(res, mat[i][0])
		}

	}

	return res
}
