package bst

import "testing"

func Test_Bst_Delete(t *testing.T) {

	bst := getBST()
	bst.Delete(60)

	bst2 := getBST()
	bst2.Delete(62)

}

func getBST() BST[int] {
	res := BST[int]{}

	res.Insert(50)
	res.Insert(30)
	res.Insert(70)
	res.Insert(40)
	res.Insert(60)
	res.Insert(80)
	res.Insert(71)
	res.Insert(200)
	res.Insert(53)
	res.Insert(66)
	res.Insert(64)
	res.Insert(67)
	res.Insert(52)
	res.Insert(55)
	res.Insert(62)
	res.Insert(65)
	return res
}
