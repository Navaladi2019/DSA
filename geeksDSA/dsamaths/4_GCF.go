package dsamaths

func findGCFConventional(a int, b int) int {

	GCF := 1

	for i := min(a, b); i > 1; i-- {

		if a%i == 0 && b%i == 0 {
			GCF = i
			break
		}
	}

	return GCF
}

func findGCFThroghEucledian(a int, b int) int {

	gcf := 1

	if b == 0 {
		gcf = a
		return gcf
	}

	return findGCFThroghEucledian(b, (max(a, b) - min(a, b)))

}

func findGCFThroghEucledianOptimized(a int, b int) int {
	if b == 0 {
		return a
	}
	return findGCFThroghEucledianOptimized(b, a%b)
}
