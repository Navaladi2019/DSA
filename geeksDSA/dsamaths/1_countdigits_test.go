package dsamaths

import (
	"testing"
)

func Test_countdigits(t *testing.T) {
	got := countDigits_ByConvertingToString(1234567891011121314)

	if 19 != got {
		t.Errorf("Expected %d but got %d", 5, got)
	}

	got = countDigits_Divideby10(1234567891011121314)

	if 19 != got {
		t.Errorf("Expected %d but got %d", 5, got)
	}
}
