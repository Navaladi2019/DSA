package algoriothmanalysis

func add(a, b int) int {
	return a + b
}

func findSumOfNaturalNumbers(ranges int) int {
	sum := 0

	for i := 0; i <= ranges; i++ {
		sum += i
	}

	return sum

}
