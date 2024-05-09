package sorting

type FnIsArrayType func(i int) bool

// this problem is similar to move zeros to left
func SortTwoTypesArray(arr []int, istype FnIsArrayType) {

	left := 0
	right := len(arr) - 1

	for {

		for {
			if !istype(arr[left]) {
				break
			}
			left++
		}

		for {
			if istype(arr[right]) {
				break
			}
			right--
		}

		if left >= right {
			break
		}

		arr[left], arr[right] = arr[right], arr[left]
	}

}

func isArrayTypeZero(i int) bool {
	if i == 0 {
		return true
	}
	return false
}

func isArrayTypeNegative(i int) bool {
	if i < 0 {
		return true
	}
	return false
}

func isArrayTypeOdd(i int) bool {

	if i < 0 {
		i = i * -1
	}

	if i%2 != 0 {
		return true
	}
	return false
}
