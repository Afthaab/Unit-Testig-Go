package calculator_test

import (
	"testing"

	"afthab.com/calculator"
)

var testcases = []struct {
	name     string
	divident int
	divisor  int
	expected int
}{
	{"division", 10, 5, 2},
	{"division_negative", -50, 10, -5},
}

func TestDivide(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Divide(tc.divident, tc.divisor)
			if got != tc.expected {
				t.Errorf("Expected %d Got %d", tc.expected, got)
			}
		})
	}

}
