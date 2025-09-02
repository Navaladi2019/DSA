package stack

func FindStockSpan(arr []int, n int) int {

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
		return n + 1
	} else {
		val, _ := sk.Pop()

		return n - val
	}

}

// // since in the above code i pop if the element is less than nth element,
//
//	it was so efficient i do not need the loop as it will only keep greater element in the stack
//
// here event instead of stack i can just use one int that will point to greater element
func FindStockSpan_1(arr []int, nthElement int) int {

	sk := ArrStack[int]{}
	if nthElement > 0 {
		sk.Push(0)
	}

	for i := 1; i < len(arr) && i < nthElement; i++ {

		for !sk.IsEmpty() {

			if val, ok := sk.Peek(); ok && arr[val] < arr[i] {
				sk.Pop()
			} else {
				break
			}
		}

		sk.Push(i)
	}

	if sk.IsEmpty() {
		return nthElement + 1
	}

	res := nthElement + 1
	for !sk.IsEmpty() {

		if v, ok := sk.Pop(); ok && arr[v] > arr[nthElement] {
			return nthElement - v
		}
	}

	return res
}

// this is a good solution if need to calculate spn for just one element but for all the elements stack is the only way
func FindStockSpanWithoutStack(arr []int, nthElement int) int {

	if nthElement == 0 {
		return 1
	}

	lastGreaterElement := -1

	for i := 0; i < nthElement; i++ {

		if arr[i] > arr[nthElement] {
			lastGreaterElement = i
		}
	}
	return nthElement - lastGreaterElement
}
