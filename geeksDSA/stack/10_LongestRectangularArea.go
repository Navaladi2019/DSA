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

func GetlargestRectangulatrArea_1(arr []int) int {
	res := 0

	sta := ArrStack[int]{}

	for i := 0; i < len(arr); i++ {
		curr := arr[i]
		for val, ok := sta.Peek(); ok && arr[val] > arr[i]; val, ok = sta.Peek() {
			sta.Pop()
			curr = max(curr, arr[val]*(i-val))
			if sta.IsEmpty() {
				curr = max(curr, arr[val]*(i))
				curr = max(curr, arr[i]*(i+1))
			} else {
				curr = max(curr, arr[i]*(i-val))
			}
		}
		res = max(res, curr)
		sta.Push(i)

	}

	for !sta.IsEmpty() {
		val, _ := sta.Pop()
		res = max(res, arr[val]*(len(arr)-(val+1)))
	}
	return res
}
