package heap

import (
	"container/heap"
	"fmt"
	"math"
)

type minDiff struct {
	diff  int
	index int
}

type minDiffItem []minDiff

func (item minDiffItem) Len() int { return len(item) }

func (item minDiffItem) Less(i, j int) bool {
	return item[i].diff < item[j].diff
}

func (pq minDiffItem) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (item *minDiffItem) Push(x any) {
	val := x.(minDiff)
	*item = append(*item, val)
}

func (item *minDiffItem) Pop() any {
	old := *item
	n := len(old)
	ele := old[n-1]
	*item = old[0 : n-1]
	return ele
}

func FindKClosestElemet(arr []int, k int, x int) {

	h := minDiffItem{}
	heap.Init(&h)
	for i, val := range arr {
		dif := float64(x - val)
		abs := math.Abs(dif)
		item := minDiff{int(abs), i}
		heap.Push(&h, item)
	}

	for i := 0; i < k; i++ {
		index := heap.Pop(&h)
		fmt.Println(arr[index.(minDiff).index])
	}
}
