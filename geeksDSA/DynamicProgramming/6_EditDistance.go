package dynamicprogramming

func EditDistanceToGetSameString(s1 string, s2 string, i int, j int) int {

	if i == 0 {
		return j
	}

	if j == 0 {
		return i
	}

	if s1[i] == s2[j] {
		return EditDistanceToGetSameString(s1, s2, i-1, j-1)
	} else {
		return 1 + min(
			//Replace
			EditDistanceToGetSameString(s1, s2, i-1, j-1),
			//Delete
			EditDistanceToGetSameString(s1, s2, i-1, j),
			//Insert
			EditDistanceToGetSameString(s1, s2, i, j-1),
		)
	}
}
