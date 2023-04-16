package main

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input int
		expected int
	} {
		{0,1},
		{1,1},
		{2,2},
		{4,24},
		{3,6},
		{5,120},
	}

	// Then we iterate over the test cases

	for _, tc := range testCases {
		result := factorial(tc.input)

		if result != tc.expected {
			t.Errorf("Factorial(%d) = %d, expected %d", tc.input, result, tc.expected)
		}
		t.Logf("Input: %d, expected: %d, result %d", tc.input, tc.expected, result)
	}
}

