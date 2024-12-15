package trie

func CountDistinctRowsInMatrix(arr [][]int) int {

	parent := NodeBin{}

	result := 0

	for i := 0; i < len(arr); i++ {
		curr := &parent
		for j := 0; j < len(arr[i]); j++ {
			if curr.child == nil {
				curr.child = make([]*NodeBin, 2)
			}
			if curr.child[arr[i][j]] == nil {
				curr.child[arr[i][j]] = &NodeBin{}
			}
			curr.child[arr[i][j]].value = arr[i][j]

			curr = curr.child[arr[i][j]]
		}

		if curr.isLeaf == false {
			result++
		}
		curr.isLeaf = true
	}

	return result
}

type NodeBin struct {
	value  int
	isLeaf bool
	child  []*NodeBin
}
