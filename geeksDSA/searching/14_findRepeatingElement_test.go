package searching

import "testing"

type RepeatingElementTc struct {
	arr  []int
	want int
}

func Test_RepeatingElement(t *testing.T) {

	got := FindRepeatingElement_Efficient([]int{1, 2, 3, 0, 3, 4, 5})

	if got != 3 {
		t.Error("has error")
	}

	got = FindRepeatingElement_Efficient([]int{0, 2, 1, 3, 2, 2})

	if got != 2 {
		t.Error("has error")
	}

	got = FindRepeatingElement_Efficient([]int{0, 0})

	if got != 0 {
		t.Error("has error")
	}
}
