package greedy

import "container/heap"

func HuffManCodingBuildTree(str string) {

	dict := make(map[rune]int, 100)

	for _, v := range []rune(str) {
		dict[v]++
	}

	nodes := make([]HuffManNode, len(dict))
	item := HuffManNodeHeap(nodes)
	heap.Init(&item)

	for k, v := range dict {
		heap.Push(&item, HuffManNode{char: k, count: v})
		for len(item) > 1 {
			left := heap.Pop(&item).(HuffManNode)
			right := heap.Pop(&item).(HuffManNode)
			heap.Push(&item, HuffManNode{char: '$', count: left.count + right.count, left: &left, right: &right})
		}
	}

	// tree := heap.Pop(&item).(HuffManNode)
	// dictCode := make(map[rune]string, 100)

	// for k, v := range dict {

	// }

}

type HuffManNode struct {
	char  rune
	count int
	left  *HuffManNode
	right *HuffManNode
}

type HuffManNodeHeap []HuffManNode

func (h HuffManNodeHeap) Len() int           { return len(h) }
func (h HuffManNodeHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h HuffManNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (item *HuffManNodeHeap) Push(x any) {
	val := x.(HuffManNode)
	*item = append(*item, val)
}

func (item *HuffManNodeHeap) Pop() any {
	old := *item
	n := len(old)
	ele := old[n-1]
	*item = old[0 : n-1]
	return ele
}
