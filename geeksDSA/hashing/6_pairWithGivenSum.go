package hashing

func FindPairWithGivenSum(arr []int, sum int) bool {

	dict := make(map[int]struct{})

	for i := 0; i < len(arr); i++ {

		if _, ok := dict[arr[i]-sum]; ok {
			return true
		} else {
			dict[arr[i]] = struct{}{}
		}
	}

	return false
}
