package arrays

import "testing"

type StockBuySelltc struct {
	ip   []int
	want int
}

var StockBuySelltcs = func() []StockBuySelltc {

	return []StockBuySelltc{
		{[]int{1, 5, 3, 8, 12}, 13},
		{[]int{30, 20, 10}, 0},
		{[]int{1, 5, 3, 1, 2, 8}, 11},
		{[]int{10, 20, 30}, 20},
	}
}

func Test_StockBuySellTestCase(t *testing.T) {

	for _, tc := range StockBuySelltcs() {

		got := findMaxProfitinStocksBuyandSell(tc.ip)
		if got != tc.want {
			t.Error("has error")
		}
	}
}
