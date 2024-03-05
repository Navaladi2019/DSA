package dsamaths

import (
	"fmt"
	"testing"
)

var testCase = func() []struct {
	input int
	want  int
} {
	return []struct {
		input int
		want  int
	}{{1, 1}, {0, 1}, {4, 24}, {5, 120}, {6, 720}, {20, 2432902008176640000}}
}

func Test_Factorial(t *testing.T) {

	for _, tc := range testCase() {
		t.Run(fmt.Sprintf("%d ! = %d", tc.input, tc.want), func(t *testing.T) {
			got := getFactorialRecursion(tc.input)

			if tc.want != got {
				t.Errorf("want %d but got %d for %d !", tc.want, got, tc.input)
			}
		})
	}

	for _, tc := range testCase() {
		t.Run(fmt.Sprintf("%d ! = %d", tc.input, tc.want), func(t *testing.T) {
			got := getFactorial(tc.input)

			if tc.want != got {
				t.Errorf("want %d but got %d for %d !", tc.want, got, tc.input)
			}
		})
	}

}

var testCaseTrailingZeros = func() []struct {
	input int
	want  int
} {
	return []struct {
		input int
		want  int
	}{{5, 1}, {10, 2}, {20, 4}}
}

var testCaseTrailingZerosLarge = func() []struct {
	input int
	want  int
} {
	return []struct {
		input int
		want  int
	}{{5, 1}, {10, 2}, {20, 4}, {25, 6}, {30, 7}, {100, 24}}
}

func Test_FactorialTrailingZeros(t *testing.T) {
	for _, tc := range testCaseTrailingZeros() {
		t.Run(fmt.Sprintf("%d ! = %d", tc.input, tc.want), func(t *testing.T) {
			got := getTrailingZerosForFactorialByConventionalMethod(tc.input)

			if tc.want != got {
				t.Errorf("want %d but got %d for %d !", tc.want, got, tc.input)
			}
		})
	}

	for _, tc := range testCaseTrailingZerosLarge() {
		t.Run(fmt.Sprintf("%d ! = %d", tc.input, tc.want), func(t *testing.T) {
			got := getTrailingZerosForFactorialByAlgorithm(tc.input)

			if tc.want != got {
				t.Errorf("want %d but got %d for %d !", tc.want, got, tc.input)
			}
		})
	}

}
