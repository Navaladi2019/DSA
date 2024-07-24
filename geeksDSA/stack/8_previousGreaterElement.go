package stack

/* stack will be usefull only when we need to find previous greater of all elements*/
func FindPreviousGreater(arr []int, n int) int {

	sk := ArrStack[int]{}
	if n > 0 {
		sk.Push(0)
	}
	for i := 1; i <= n; i++ {

		for !sk.IsEmpty() {
			valIndex, _ := sk.Peek()
			if arr[valIndex] < arr[n] {
				sk.Pop()
			} else {
				break
			}
		}
		if i != n {
			sk.Push(i)
		}
	}

	if sk.IsEmpty() {
		return -1
	} else {
		val, _ := sk.Pop()

		return arr[val]
	}

}

func FindPreviousGreater_myEfficient(arr []int, n int) int {

	for i := n - 1; i >= 0; i-- {

		if arr[i] > arr[n] {
			return arr[i]
		}

	}

	return -1
}
