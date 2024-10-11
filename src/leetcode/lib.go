package main

import (
    "strconv"
)

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