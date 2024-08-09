package heap

func BuyingMaxItemWIthGivenSuym(arr []int, sum int) int {

	count := 0
	q := MinHeap{}
	q.Init(arr)
	for sum > 0 {
		if ok, val := q.Extract(); ok {
			if sum-val >= 0 {
				count++
			}
		} else {
			break
		}
	}

	return count
}
