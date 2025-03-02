package greedy

var count int = 0

func GetMinCoins(arr []int, sum int) int {

	for i := 0; i < len(arr); i++ {
		if sum-arr[i] < 0 {
			continue
		}
		sum = sum - arr[i]
		count++
		GetMinCoins(arr, sum)
	}

	return count
}
