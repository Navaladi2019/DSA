package dynamicprogramming

import (
	"testing"
)

func Test_LCS(t *testing.T) {

	got := LCS("AXYZ", "BAZ", 3, 2)

	if got != 2 {
		t.Error("Has error n naive method")
	}

	got = FindLCSUsinMemoization("AXYZ", "BAZ")

	if got != 2 {
		t.Error("Has error n memo method")
	}

	got = FindLCSUsinTabulation("AXYZ", "BAZ")

	if got != 2 {
		t.Error("Has error n table method")
	}

}
