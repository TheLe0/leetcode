package main

import (
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