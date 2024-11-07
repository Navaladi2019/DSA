package backtracking

func CanRatReachCheese(arr [][]int, i int, j int) bool {

	if i == len(arr)-1 && j == len(arr[0])-1 {
		return true
	}
	if !CanRatReachCheeseIsSafe(arr, i, j) {
		return false
	}

	if CanRatReachCheese(arr, i+1, j) {
		return true
	}
	return CanRatReachCheese(arr, i, j+1)

}

func CanRatReachCheeseIsSafe(arr [][]int, i int, j int) bool {

	if i >= len(arr) || j >= len(arr[0]) || arr[i][j] == 0 {
		return false
	}

	return true
}
