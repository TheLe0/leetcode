package main

import (
	"strconv"
)

// 1. Two Sum
// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for idx, value := range nums {

		rest := target - value

		if index, exists := numsMap[rest]; exists {
			return []int{index, idx}
		}

		numsMap[value] = idx
	}

	return nil
}

// 738. Monotone Increasing Digits
// https://leetcode.com/problems/monotone-increasing-digits/description/

func validateDigits(n int) bool {
	numStr := strconv.Itoa(n)

	for i := 0; i < len(numStr)-1; i++ {
		currentDigit := numStr[i]
		nextDigit := numStr[i+1]

		if currentDigit > nextDigit {
			return false
		}
	}

	return true
}

func monotoneIncreasingDigits(n int) int {
	goNext := true
	monotoneValue := n

	for goNext {
		isMonotone := validateDigits(monotoneValue)

		if isMonotone {
			goNext = false
		} else {
			monotoneValue--
		}
	}

	return monotoneValue
}
