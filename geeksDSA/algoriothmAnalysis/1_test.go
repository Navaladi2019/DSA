package algoriothmanalysis

import (
	"testing"
)

func Test_add(t *testing.T) {
	got := add(4, 5)

	if 9 != got {
		t.Errorf("Expected %d but got %d", 8, got)
	}
}

func Test_findSumOfNaturalNumbers(t *testing.T) {
	got := findSumOfNaturalNumbers(5)

	if 15 != got {
		t.Errorf("Expected %d but got %d", 15, got)
	}

}
