package hashing

// here we insert only after checking if the pair sum sodes not exist
// [1,3,2,4]. If we insert everything in one go and the it will look for 3 in loop giving 3=3 =6 which is not right
// so we check if arr[i]- sum exist in dict, if it does not then we insert arr[i] into dictionary

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
