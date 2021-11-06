package main

import (
	"log"
	"unicode"
)

func lengthOfLastWord(s string) int {
	currLen := 0
	foundStartPos := false

	for i := len(s) - 1; i >= 0; i-- {
		isSpace := unicode.IsSpace(rune(s[i]))
		if isSpace && !foundStartPos {
			continue
		} else if isSpace && foundStartPos {
			return currLen
		} else {
			foundStartPos = true
			currLen++
		}
	}

	return currLen
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "Hello World", expected: 5},
		{input: "   fly me   to   the moon  ", expected: 4},
		{input: "luffy is still joyboy", expected: 6},
	}

	for _, test := range testCases {
		got := lengthOfLastWord(test.input)
		if got != test.expected {
			log.Fatalf(
				"expected %d with an input of %s, but got %d",
				test.expected,
				test.input,
				got,
			)
		}
	}
}
