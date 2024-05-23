package stack

func FindStockSpan(arr []int, n int) int {

	sk := ArrStack[int]{}

	if n > 0 {
		sk.Push(0)
	}

	for i := 1; i <= n; i++ {

		for !sk.isEmpty() {
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

	if sk.isEmpty() {
		return n + 1
	} else {
		val, _ := sk.Pop()

		return n - val
	}

}
