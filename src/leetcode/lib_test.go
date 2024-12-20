package main

import (
	"reflect"
	"testing"
)

func TestMonotoneIncreasingDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 9},
		{1234, 1234},
		{332, 299},
	}

	for _, test := range tests {
		result := monotoneIncreasingDigits(test.input)
		if result != test.expected {
			t.Errorf("monotoneIncreasingDigits(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums=%v and target=%d, expected %v, but got %v",
				test.nums, test.target, test.expected, result)
		}
	}
}
