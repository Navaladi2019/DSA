package arrays

func findMaxProfitinStocksBuyandSell(arr []int) int {

	profit := 0
	minvalue := arr[0]
	maxvalue := arr[0]

	for i := 1; i < len(arr); i++ {

		if arr[i] < maxvalue {
			profit = profit + (maxvalue - minvalue)
			minvalue = arr[i]
			maxvalue = arr[i]
		}

		if arr[i] > maxvalue {
			maxvalue = arr[i]
		}
	}

	profit = profit + (maxvalue - minvalue)

	return profit
}

func FindStockProfit(arr []int) int {
	profit := 0

	lastLowVal := arr[0]

	for i := 0; i < len(arr); i++ {

		if arr[i] < lastLowVal {
			lastLowVal = arr[i]
		}

		if arr[i] > lastLowVal {
			profit = profit + arr[i] - lastLowVal
			lastLowVal = arr[i]
		}
	}

	return profit
}
