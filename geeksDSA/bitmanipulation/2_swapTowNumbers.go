package bitmanipulation

func SwapTwoNumbers(a int, b int) (int, int) {

	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func CheckIfIthBitIsSetOrNot(n int, i int) (bool, bool) {

	right := (n>>i)&1 != 0
	left := (1<<i)&n != 0

	return left, right

}
