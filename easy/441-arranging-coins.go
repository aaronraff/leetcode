package main

import "log"

func arrangeCoins(n int) int {
	step := 0

	// Can we fill the next step?
	for n >= step + 1 {
		step++
		n -= step
	}

	return step
}

func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 5, expected: 2},
		{input: 8, expected: 3},
	}

	for _, test := range testCases {
		got := arrangeCoins(test.input)
		if got != test.expected {
			log.Fatalf(
				"expected %d with an input of %d, but got %d",
				test.expected,
				test.input,
				got,
			)
		}
	}
}
