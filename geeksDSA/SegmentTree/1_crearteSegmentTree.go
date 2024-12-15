package segmenttree

func CreateSegmentTree(tree []int, arr []int, start int, end int, SegmentIndex int) int {

	if start == end {

		tree[SegmentIndex] = arr[start]
		return arr[start]
	}

	mid := (start + end) / 2

	tree[SegmentIndex] = CreateSegmentTree(tree, arr, start, mid, 2*SegmentIndex+1) +
		CreateSegmentTree(tree, arr, mid+1, end, 2*SegmentIndex+2)

	return tree[SegmentIndex]
}

func GetRangeInTree(tree []int, qs int, qe int, siStart int, siEnd int, si int) int {

	if siEnd < qs || siStart > qe {
		return 0
	}

	if siStart >= qs && siEnd <= qe {
		return tree[si]
	}
	mid := (siStart + siEnd) / 2
	return GetRangeInTree(tree, qs, qe, siStart, mid, 2*si+1) + GetRangeInTree(tree, qs, qe, mid+1, siEnd, 2*si+2)
}

func UpdateInTree(tree []int, index int, val int, siStart int, siEnd int, si int) int {

	if index > siEnd || index < siStart {
		return tree[si]
	}

	if siStart == siEnd && siStart == index {
		tree[si] = val
		return val
	}

	mid := (siStart + siEnd) / 2

	tree[si] = UpdateInTree(tree, index, val, siStart, mid, 2*si+1) + UpdateInTree(tree, index, val, mid+1, siEnd, 2*si+2)

	return tree[si]

}

func UpdateInTreeWithDif(tree []int, index int, diff int, siStart int, siEnd int, si int) {

	if index > siEnd || index < siStart {
		return
	}

	tree[si] = tree[si] + diff

	if siStart == siEnd {
		return
	}

	mid := (siStart + siEnd) / 2

	UpdateInTreeWithDif(tree, index, diff, siStart, mid, 2*si+1)
	UpdateInTreeWithDif(tree, index, diff, mid+1, siEnd, 2*si+2)

}
