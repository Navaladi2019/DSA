package dynamicprogramming

func MaxINcreasingSubSequenceNaive(arr []int) int {

	res := 0

	for i := 0; i < len(arr); i++ {
		loopres := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				loopres += arr[j]
			}

		}
		res = max(res, loopres)
	}
	return res
}

func MaxINcreasingSubSequenceDynamic(arr []int) int {

	ls := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		val := 0
		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				val = max(ls[j], val)
			}
		}
		ls[i] = arr[i] + val
	}

	res := 0
	for i := 0; i < len(ls); i++ {
		res = max(res, ls[i])
	}
	return res
}
