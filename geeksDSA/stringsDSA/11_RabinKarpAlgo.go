package stringsDSA

func RBKSearch(str string, patt string) (bool, []int) {

	haspattern := false
	patternIndexdx := make([]int, 0, 10)

	d := 256
	q := 101
	h := 1

	// Compute (d^(M-1))%q
	for i := 0; i < len(patt); i++ {
		h = (h * d) % q
	}

	strByte := []byte(str)
	patByte := []byte(patt)

	// caluctating hash for index 0  and pattern
	// here we are taking modulus of q so that integer overflow will not happen
	p, t := 0, 0
	for i := 0; i < len(patByte); i++ {
		p = (p * d) + int(patByte[i])%q
		t = (t * d) + int(strByte[i])%q
	}

	for i := 0; i <= len(str)-len(patt); i++ {

		if p == t {
			j := 0
			for j < len(patt) {

				if str[i+j] != patt[i] {
					break
				}
				j++
			}
			if j == len(patt) {
				haspattern = true
				patternIndexdx = append(patternIndexdx, i)
			}
		}

		if i < len(str)-len(patt) {
			//here the formula is d * (h[i]-(txt[i]*(d^m-1)))+txt[i+m]
			t = (d*(t-int(strByte[i])*h) + int(strByte[i+len(patt)])) % q
		}

	}

	return haspattern, patternIndexdx
}
