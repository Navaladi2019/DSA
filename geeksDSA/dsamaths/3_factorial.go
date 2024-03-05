package dsamaths

func getFactorial(input int) int {

	if input <= 1 {
		return 1
	}
	op := 1
	for i := 2; i <= input; i++ {
		op = op * i
	}

	return op
}

/* it has theta(n) fopr space as function calls are stored in stack*/
func getFactorialRecursion(i int) int {
	if i <= 1 {
		return 1
	}
	return i * getFactorialRecursion(i-1)
}

/*This causes interger overflow for number of 100!*/
func getTrailingZerosForFactorialByConventionalMethod(ip int) int {

	fact := getFactorial(ip)

	trailingZeros := 0

	for fact%10 == 0 && fact > 0 {
		trailingZeros++
		fact = fact / 10
	}

	return trailingZeros
}

/*
we are going to use a logic called prime factorial of 5\
5-> 5*1 so one trailing zero
10->5*2 so 2 zeros
15 -> 5*3 so 3 zeros
20 -> 5*4 so 4 zeros
25 -> 5*5 here we have two 5 of multiplication so we have to add 1 -> (5*4+(5*1+5*1)) -> 6 Zeros

thats why we use prime factoruial of 5 -> 5,25,125,....

so formula would be => (n/5)+(n/25)+(n/125)..... till n > divisor
*/
func getTrailingZerosForFactorialByAlgorithm(ip int) int {

	if ip < 5 {
		return 0
	}
	res := 0

	for i := 5; i <= ip; i = i * 5 {
		res = res + ip/i
	}

	return res
}
