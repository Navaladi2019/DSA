package stack

func GetLargestRectangularArea(arr []int) int {

	st := ArrStack[int]{}

	res := 0

	for i := 0; i < len(arr); i++ {

		for val, ok := st.Peek(); ok && arr[val] >= arr[i]; val, ok = st.Peek() {
			st.Pop()
			curr := arr[val]
			if st.IsEmpty() {
				curr = curr * i
			} else {
				nexttop, _ := st.Peek()

				curr = curr * (i - nexttop - 1)
			}

			res = max(res, curr)

		}
		st.Push(i)
	}

	for !st.IsEmpty() {

		val, _ := st.Pop()

		curr := arr[val]

		if st.IsEmpty() {
			curr = curr * len(arr)
		} else {
			nexttop, _ := st.Peek()
			curr = curr * (len(arr) - nexttop - 1)
		}
		res = max(res, curr)
	}
	return res
}
