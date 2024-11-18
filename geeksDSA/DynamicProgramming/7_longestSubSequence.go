package dynamicprogramming

// conside tyhe array [3,4,2,8,10,5,1] for ths the longes sub sequence is [3,4,8,10]
// because here by having 3 we do not consider numbers less than tha

//[3],[3,4],[3,4,8],[3,4,8,10],[3,5],[3,8,10],[3,8],[3,10] so for each index we will be computing the subsequence

func LongestSubsequence(arr []int) int {

	ls := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		val := 0

		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[i] {
				val = max(val, ls[j])
			}
		}
		ls[i] = 1 + val
	}
	res := ls[0]

	for i := 1; i < len(ls); i++ {
		res = max(res, ls[i])
	}

	return res

}

func LongestIncreasingSubSequence(arr []int) int {
	tail := make([]int, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		ceilIndex := findCeilIndex(tail, arr[i])
		if ceilIndex == -1 {
			tail = append(tail, arr[i])
		} else {
			tail[ceilIndex] = arr[i]
		}
	}
	return len(tail)
}

func findCeilIndex(arr []int, val int) int {

	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1
	result := -1

	for low <= high {

		mid := (low + high) / 2

		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result

}
