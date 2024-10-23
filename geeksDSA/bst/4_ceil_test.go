package bst

import "testing"

func Test_Ceil(t *testing.T) {

	bst := getBST()

	got := Ceil_1(bst.root, 59)

	if *got != 60 {
		t.Error("has error in 59")
	}

	got = Ceil_1(bst.root, 201)

	if got != nil {
		t.Error("has error in 201")
	}

	got = Ceil_1(bst.root, 66)

	if *got != 66 {
		t.Error("has error in 66")
	}

	got = Ceil_1(bst.root, 63)

	if *got != 64 {
		t.Error("has error in 63")
	}

}
