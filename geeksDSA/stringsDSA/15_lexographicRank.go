package stringsDSA

/*

BAC -. we sort it and find permutation and tell at which count the target string is formed0

*/
func LexographicRank(str string) int {

	count := 0
	res := 0
	for i := 0; i < len(str); i++ {
		count = 0
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				count++
			}
		}

		factt := fact(len(str) - i - 1)
		res = res + (count * factt)
	}
	return res + 1
}

func fact(i int) int {

	if i <= 1 {
		return 1
	}

	return i * fact(i-1)
}
