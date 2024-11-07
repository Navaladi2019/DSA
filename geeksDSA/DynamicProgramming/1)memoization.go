package dynamicprogramming

var fibMap = make(map[int]int)

func Fib(n int) int {

	val, ok := fibMap[n]
	if ok {
		return val
	}
	if n == 0 || n == 1 {
		return n
	}
	val = Fib(n-1) + Fib(n-2)
	fibMap[n] = val
	return val
}
