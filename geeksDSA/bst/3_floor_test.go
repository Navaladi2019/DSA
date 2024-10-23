package bst

import "testing"

func Test_Floor_iterative(t *testing.T) {

	bst := getBST()

	got := Floor_iterative1(bst.root, 63)

	if *got != 62 {
		t.Error("has error in 63")
	}

	got = Floor_iterative1(bst.root, 54)

	if *got != 53 {
		t.Error("has error in 54")
	}

	got = Floor_iterative1(bst.root, 58)

	if *got != 55 {
		t.Error("has error in 58")
	}

	got = Floor_iterative1(bst.root, 1)

	if got != nil {
		t.Error("has error in 1")
	}
}
