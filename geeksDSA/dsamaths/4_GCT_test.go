package dsamaths

import (
	"fmt"
	"testing"
)

var tcs = func() []struct {
	a    int
	b    int
	want int
} {

	return []struct {
		a    int
		b    int
		want int
	}{
		{5, 15, 5},
		{2, 4, 2},
		{3, 9, 3},
		{200, 150, 50},
		{7, 9, 1},
	}
}

func Test_FindGCFConventional(t *testing.T) {

	for _, tc := range tcs() {

		t.Run(fmt.Sprintf("%d & %d = %d", tc.a, tc.b, tc.want), func(t *testing.T) {

			got := findGCFConventional(tc.a, tc.b)

			if got != tc.want {
				t.Errorf("Expected GCF to be %d for %d & %d but got %d", tc.want, tc.a, tc.b, got)
			}
		})
	}

}

func Test_FindGCFEucledian(t *testing.T) {

	for _, tc := range tcs() {

		t.Run(fmt.Sprintf("%d & %d = %d", tc.a, tc.b, tc.want), func(t *testing.T) {

			got := findGCFThroghEucledian(max(tc.a, tc.b), min(tc.a, tc.b))

			if got != tc.want {
				t.Errorf("Expected GCF to be %d for %d & %d but got %d", tc.want, tc.a, tc.b, got)
			}
		})
	}

}

func Test_findGCFThroghEucledianOptimized(t *testing.T) {

	for _, tc := range tcs() {

		t.Run(fmt.Sprintf("%d & %d = %d", tc.a, tc.b, tc.want), func(t *testing.T) {

			got := findGCFThroghEucledianOptimized(max(tc.a, tc.b), min(tc.a, tc.b))

			if got != tc.want {
				t.Errorf("Expected GCF to be %d for %d & %d but got %d", tc.want, tc.a, tc.b, got)
			}
		})
	}

}
