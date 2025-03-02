package dynamicprogramming

var dict = make(map[int]int)

func GetFib(n int) int {

	if n == 1 || n == 0 {
		return n
	}
	if val, ok := dict[n]; ok {
		return val
	}

	dict[n] = GetFib(n-1) + GetFib(n-2)
	return dict[n]
}
