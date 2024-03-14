package arrays

import "testing"

type RemoveDuplicateArraytc struct {
	ip   []int
	want RemoveDuplicateArrayWant
}

type RemoveDuplicateArrayWant struct {
	arr []int
	len int
}

var RemoveDuplicateArraytcs = func() []RemoveDuplicateArraytc {
	return []RemoveDuplicateArraytc{
		{[]int{10, 20, 20, 30, 30, 30, 30}, RemoveDuplicateArrayWant{arr: []int{10, 20, 30, 30, 30, 30, 30}, len: 3}},
		{[]int{10, 10, 10}, RemoveDuplicateArrayWant{arr: []int{10, 10, 10}, len: 1}},
	}
}

func Test_CheckRemovingDuplicateFromArray(t *testing.T) {

	for _, tc := range RemoveDuplicateArraytcs() {
		l := RemoveDuplicateFromAnArray(tc.ip)

		if l != tc.want.len {
			t.Error("error")
		}

		for i := 0; i < len(tc.want.arr); i++ {

			if tc.ip[i] != tc.want.arr[i] {
				t.Error("err")
				break
			}
		}

	}
}
