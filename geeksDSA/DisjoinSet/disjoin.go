package disjoinset

func InitializeDisjoinSets(arr []int) {

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

}

func Union(i int, j int, arr []int) {

	iParent := Find(i, arr)
	jParent := Find(j, arr)

	if iParent == jParent {
		return
	}

	arr[jParent] = iParent

}

func Find(x int, arr []int) int {

	if arr[x] == x {
		return x
	}
	return Find(arr[x], arr)
}

func IsFriends(i int, j int, arr []int) bool {

	return Find(i, arr) == Find(j, arr)
}
