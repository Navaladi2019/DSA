package sorting

//TODO: i have to watch this video again
/* on^2*/
func FindInversion(arr []int) int {

	res := 0
	for i := 0; i < len(arr); i++ {

		for j := i; j < len(arr); j++ {

			if arr[i] > arr[j] {
				res++
			}
		}
	}

	return res
}

func CountInversion(arr []int, left int, right int) int {

	result := 0

	if left < right {

		mid := left + (right-left)/2

		result += CountInversion(arr, left, mid)
		result += CountInversion(arr, mid+1, right)
		result += CountAndMerge(arr, left, mid, right)
	}
	return result
}

func CountAndMerge(arr []int, low int, mid int, high int) int {

	result := 0
	leftArr := make([]int, mid-low+1)
	rightArr := make([]int, high-mid)

	copy(leftArr, arr[low:mid+1])
	copy(rightArr, arr[mid+1:high+1])

	i := 0

	j := 0
	k := low

	for i < len(leftArr) && j < len(rightArr) {

		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
			k++
		} else {
			arr[k] = rightArr[j]

			// now arr[j] is smaller that all the arr2[i] right side so that forms inversion
			result += len(leftArr) - i
			j++
			k++
		}
	}

	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}

	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}

	return result
}
