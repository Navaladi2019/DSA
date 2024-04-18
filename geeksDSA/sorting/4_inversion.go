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
