package dsamaths

func fincLCMCONventional(a int, b int) int {

	target := max(a, b)

	for target%a != 0 || target%b != 0 {
		target = target + max(a, b)
	}

	return target
}

/*
a*b = gcd(a,b)* lcm(a,b) =>
lcm(a,b) = (a*b)/gcd(a,b)
*/
func fincLCMOptimized(a int, b int) int {
	return (a * b) / findGCFThroghEucledian(a, b)

}
