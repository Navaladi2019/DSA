package arrays

func findMaximunDifferenceArray(arr []int) int {

	maxdifference := arr[1] - arr[0]
	lowestLeftNumber := arr[0]

	for i := 1; i < len(arr); i++ {
		maxdifference = max(maxdifference, arr[i]-lowestLeftNumber)
		lowestLeftNumber = min(lowestLeftNumber, arr[i])
	}

	return maxdifference
}

func FinDmAximunDifference(arr []int) int {

	i := 0

	diff := arr[1] - arr[0]

	for j := 1; j < len(arr); j++ {
		diff = max(diff, arr[j]-arr[i])
		if arr[j] < arr[i] {
			i = j
		}
	}

	return diff
}
