package bst

import "testing"

func Test_InsetBstIntro(t *testing.T) {

	bt := TestBst{}
	bt.Insert(50)
	bt.Insert(70)
	bt.Insert(30)
	bt.Insert(40)
	bt.Insert(10)
	bt.Insert(60)
	bt.Insert(80)

	got := bt.Search(60)

	if got != true {
		t.Error("test failed to find key 60")
	}
}
