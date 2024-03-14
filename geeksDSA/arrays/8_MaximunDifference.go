package arrays

func findMaximunDifferenceArray(arr []int) int {

	maxdifference := arr[1] - arr[0]
	lowestLeftNumber := arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i]-lowestLeftNumber > maxdifference {
			maxdifference = arr[i] - lowestLeftNumber
		}

		if lowestLeftNumber > arr[i] {
			lowestLeftNumber = arr[i]
		}

	}

	return maxdifference
}
