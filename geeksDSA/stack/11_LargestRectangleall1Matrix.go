package stack

func LargestRectange(arr [][]int) int {

	res := 0

	prevArea := make([]int, len(arr[0]))

	for i := 0; i < len(arr); i++ {
		CurrArea := make([]int, 0, len(arr[0]))
		for j := 0; j < len(prevArea); j++ {
			currvalue := arr[i][j]

			if currvalue > 0 {
				currvalue = currvalue + prevArea[j]
			}

			CurrArea = append(CurrArea, currvalue)
		}

		res = max(GetLargestRectangularArea(CurrArea), res)

		prevArea = CurrArea

	}

	return res
}
