package dsamaths

func istheNumberPalindrome(i int) bool {
	initial := i
	rev := 0
	for i > 0 {
		rev = rev*10 + i%10
		i = i / 10
	}

	return initial == rev
}

func isArrayNumbersPalindrome[T comparable](arr []T) bool {

	for i := 0; i <= len(arr) && i < len(arr); i++ {

		if arr[i] != arr[(len(arr)-1-i)] {
			return false
		}
	}

	return true
}
