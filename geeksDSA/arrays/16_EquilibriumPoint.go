package arrays

func IsArrayHasEquilibriumPoint(arr []int) bool {

	PrefixSumFronttoBack := make([]int, len(arr))
	copy(PrefixSumFronttoBack, arr)

	PrefixSumBaxckToFront := make([]int, len(arr))
	copy(PrefixSumBaxckToFront, arr)

	for i := 1; i < len(arr); i++ {

		PrefixSumFronttoBack[i] += PrefixSumFronttoBack[i-1]
		PrefixSumBaxckToFront[len(arr)-1-i] += PrefixSumBaxckToFront[len(arr)-1-i+1]
	}

	left := 0
	for i := 0; i < len(arr)-1; i++ {

		if PrefixSumBaxckToFront[i+1] == left {
			return true
		}
		left = PrefixSumFronttoBack[i]
	}

	return false

}

func IsArrayHasEquilibriumPointEfficientSolution(arr []int) bool {

	rightSum := 0

	for i := 0; i < len(arr); i++ {
		rightSum += arr[i]
	}

	leftSum := 0

	for i := 0; i < len(arr); i++ {
		rightSum -= arr[i]

		if rightSum == leftSum {
			return true
		}

		leftSum += arr[i]

	}

	return false
}
