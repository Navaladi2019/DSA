package trie

import "testing"

func Test_Trie(t *testing.T) {
	trie := Node{}
	trie.Add("naval")
	got := trie.Search("naval")

	if got == false {
		t.Error("error in trie")
	}

	got = trie.Search("navals")

	if got == true {
		t.Error("error in trie")
	}
	got = trie.Search("nval")

	if got == true {
		t.Error("error in trie")
	}
}
