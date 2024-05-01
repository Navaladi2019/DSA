package hashing

import "container/list"

type MyHash struct {
	bucket int
	table  []list.List
}

func (h *MyHash) MyHash(b int) {
	h.bucket = b
	h.table = make([]list.List, b)
	for i := range h.table {
		h.table[i].Init()
	}
}

func (h *MyHash) Insert(key int) {
	i := key % h.bucket
	h.table[i].PushBack(key)
}

func (h *MyHash) Delete(key int) {

	i := key % h.bucket

	for e := h.table[i].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			h.table[i].Remove(e)
			return
		}
	}

}

func (h *MyHash) Search(key int) bool {
	i := key % h.bucket

	for e := h.table[i].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			return true
		}
	}

	return false
}
