package segmenttree

func CreateBinaryIndexTree(arr []int, tree []int, s int, e int) {

	if s < e {

		return
	}

	if s == e {
		tree[s] = arr[s-1]
		return
	}

	for i := 0; s+(2^i) <= e; i++ {

		tree[s+(2^i)] = tree[s+(2^(i-1))] + getSumFromINdex(s+(2^(i-1)), s+(2^i), arr)

	}

}

func getSumFromINdex(s int, e int, arr []int) int {

	res := 0
	for i := s; i <= e; i++ {
		res += arr[i]
	}
	return res

}
