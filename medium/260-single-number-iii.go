package main

import (
	"log"
	"reflect"
	"sort"
)

func singleNumber(nums []int) []int {
	var single []int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		isLastElement := i == len(nums) - 1
		if !isLastElement && nums[i + 1] == nums[i] {
			// Skip over the duplicate element.
			i++
		} else {
			single = append(single, nums[i])
		}
	}

	return single
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 1, 3, 2, 5}, expected: []int{3, 5}},
		{input: []int{-1, 0}, expected: []int{-1, 0}},
		{input: []int{0, 1}, expected: []int{0, 1}},
	}

	for _, test := range testCases {
		got := singleNumber(test.input)
		if !reflect.DeepEqual(got, test.expected) {
			log.Fatalf(
				"expected %v with an input of %v, but got %v",
				test.expected,
				test.input,
				got,
			)
		}
	}
}
