package arrays

/*
{2,0,2} -> 2

{3,0,1,2,5} -> 6

*/

func FindAmountofTrappingRainWater(arr []int) int {

	amount := 0

	lmax := make([]int, len(arr))
	rmax := make([]int, len(arr))

	lmax[0] = arr[0]
	rmax[len(arr)-1] = arr[len(arr)-1]

	for i := 1; i < len(arr); i++ {

		lmax[i] = max(arr[i], lmax[i-1])
	}

	for i := len(arr) - 2; i >= 0; i-- {

		rmax[i] = max(arr[i], rmax[i+1])
	}

	for i := 0; i < len(arr); i++ {

		amount += min(lmax[i], rmax[i]) - arr[i]
	}

	return amount
}
