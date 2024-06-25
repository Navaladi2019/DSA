package arrays

import "testing"

type MaxInArrayTc struct {
	input []int
	want  int
}

var MaxInArrayTcs = func() []MaxInArrayTc {

	return []MaxInArrayTc{
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9},

		{input: []int{1, 999, 3, 100, 5, 6, 300, 8, 9},
			want: 999},
	}
}

var MaxInArraySecondTcs = func() []MaxInArrayTc {

	return []MaxInArrayTc{
		{input: []int{25, 20, 12, 20, 10},
			want: 1},
		{input: []int{10, 5, 8, 20},
			want: 0},
		{input: []int{10, 10, 10},
			want: -1},
	}
}

func Test_FindMaxinArray(t *testing.T) {

	for _, tc := range MaxInArrayTcs() {
		got := findLargestElementInArray(tc.input)
		if tc.want != got {
			t.Error("has Error")
		}
	}
}

func Test_FindSecondIndexMaxinArray(t *testing.T) {

	for _, tc := range MaxInArraySecondTcs() {
		got := FindSecondLargestElementIndex(tc.input)
		if tc.want != got {
			t.Error("has Error")
		}
	}
}
