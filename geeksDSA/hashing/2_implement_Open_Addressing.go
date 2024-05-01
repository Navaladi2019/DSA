package hashing

import "errors"

type MyHashOpenAddressing struct {
	arr  []MyHashOpenAddressingVal
	cap  int
	size int
}

type MyHashOpenAddressingVal struct {
	val       int
	hasValue  bool
	isdeleted bool
}

func (h *MyHashOpenAddressing) Hash(key int) int {
	return key % h.cap
}

func (h *MyHashOpenAddressing) Init(c int) *MyHashOpenAddressing {
	h.cap = c
	h.size = 0
	h.arr = make([]MyHashOpenAddressingVal, c)
	return h
}

func (h *MyHashOpenAddressing) Insert(key int) error {
	hash := h.Hash(key)

	i := hash

	for h.arr[i].hasValue {

		if h.arr[i].val == key {
			break
		}
		i = (i + 1) % h.cap

		if hash == i {
			break
		}
	}

	if !h.arr[i].hasValue {
		h.arr[i].val = key
		h.arr[i].isdeleted = false
		h.arr[i].hasValue = true
		h.size++
	} else if h.arr[i].hasValue && h.arr[i].val == key {
		h.arr[i].val = key
		h.arr[i].isdeleted = false
		h.arr[i].hasValue = true
	} else if h.arr[i].hasValue && h.size == h.cap {

		return errors.New("Size fully occupied")
	}

	return nil
}

func (h *MyHashOpenAddressing) Search(key int) int {

	hash := h.Hash(key)

	i := hash

	for h.arr[i].hasValue || h.arr[i].isdeleted {

		if !h.arr[i].isdeleted && h.arr[i].val == key {
			return i
		}

		i = (i + 1) % h.cap

		if hash == i {
			return -1
		}
	}

	return -1
}

func (h *MyHashOpenAddressing) Delete(key int) {
	hash := h.Hash(key)

	i := hash

	for !h.arr[i].hasValue || h.arr[i].isdeleted {

		i = (i + 1) % h.cap

		if hash == i {
			break
		}
	}

	if !h.arr[i].isdeleted && h.arr[i].hasValue && h.arr[i].val == key {
		h.arr[i].isdeleted = true
		h.arr[i].hasValue = false
		h.size--

	}
}

func (h *MyHashOpenAddressing) GetAll() []MyHashOpenAddressingVal {
	return h.arr
}
