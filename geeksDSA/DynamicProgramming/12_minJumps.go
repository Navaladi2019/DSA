package dynamicprogramming

func MinJumps(arr []int, index int) int {

	if arr[index]+index >= len(arr)-1 {
		return 1
	}
	nextIndex := index + 1
	for i := 1; i < arr[index]; i++ {
		nextIndex = max(arr[i]+i-arr[index], arr[nextIndex]+nextIndex-arr[index])
	}
	return 1 + MinJumps(arr, nextIndex)
}
