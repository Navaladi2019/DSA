package segmenttree

import "testing"

func Test_CreateSegmentTree(t *testing.T) {

	arr := []int{10, 20, 30, 40}

	tree := make([]int, len(arr)*4)

	CreateSegmentTree(tree, arr, 0, len(arr)-1, 0)

	got := GetRangeInTree(tree, 0, 2, 0, len(arr)-1, 0)

	if got != 60 {

		t.Error("has no error in segment tree.")
	}

	_ = UpdateInTree(tree, 0, 20, 0, 4, 0)

	got = GetRangeInTree(tree, 0, 2, 0, len(arr)-1, 0)

	if got != 70 {

		t.Error("has no error in segment tree.")
	}

}

func Test_UpdateSegmentTree(t *testing.T) {

	arr := []int{10, 20, 30, 40}

	tree := make([]int, len(arr)*4)

	CreateSegmentTree(tree, arr, 0, len(arr)-1, 0)

	got := GetRangeInTree(tree, 0, 2, 0, len(arr)-1, 0)

	if got != 60 {

		t.Error("has no error in segment tree.")
	}

	diff := 20 - arr[0]
	UpdateInTreeWithDif(tree, 0, diff, 0, 4, 0)

	got = GetRangeInTree(tree, 0, 2, 0, len(arr)-1, 0)

	if got != 70 {

		t.Error("has no error in segment tree.")
	}

}
