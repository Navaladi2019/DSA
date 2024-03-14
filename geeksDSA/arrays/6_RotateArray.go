package arrays

/*In Golang way this is very easy*/
func RotateArray(arr []int, n int) []int {
	n = n % (len(arr))

	res := append(arr[n:], arr[:n]...)

	return res

}

func RotateArrayLogic(arr []int, n int) []int {
	n = n % (len(arr))
	ReverseAnArray(arr)
	ReverseAnArray(arr[:len(arr)-n])
	ReverseAnArray(arr[len(arr)-n:])
	return arr
}
