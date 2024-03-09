package dsamaths

func isPrime(i int) bool {

	if i == 2 || i == 3 || i == 5 || i == 7 || i == 11 {
		return true
	} else if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%11 == 0 {
		return false
	} else {
		return true
	}
}

func isPrimeTest(i int) bool {
	if i == 1 {
		return false
	}

	if i == 2 || i == 3 {
		return true
	}

	if i%2 == 0 || i%3 == 0 {
		return false
	}

	for n := 5; n*n <= n; i = i + 6 {

		if i%n == 0 || i%(n/2) == 0 {
			return false
		}
	}

	return true
}
