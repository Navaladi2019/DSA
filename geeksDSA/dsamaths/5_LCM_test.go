package dsamaths

import (
	"fmt"
	"testing"
)

var tcs_LCM = func() []struct {
	a    int
	b    int
	want int
} {

	return []struct {
		a    int
		b    int
		want int
	}{
		{4, 6, 12},
		{12, 15, 60},
		{2, 8, 8},
		{50, 59, 2950},
	}
}

func Test_FindLCMConventional(t *testing.T) {

	for _, tc := range tcs_LCM() {

		t.Run(fmt.Sprintf("%d & %d = %d", tc.a, tc.b, tc.want), func(t *testing.T) {

			got := fincLCMCONventional(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("Expected LCM to be %d for %d & %d but got %d", tc.want, tc.a, tc.b, got)
			}
		})
	}

}

func Test_FindLCMOptimized(t *testing.T) {

	for _, tc := range tcs_LCM() {

		t.Run(fmt.Sprintf("%d & %d = %d", tc.a, tc.b, tc.want), func(t *testing.T) {

			got := fincLCMOptimized(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("Expected LCM to be %d for %d & %d but got %d", tc.want, tc.a, tc.b, got)
			}
		})
	}

}
