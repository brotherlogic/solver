package main

import (
	"math"
	"strconv"
)

func isPrime(v int64) bool {
	if v%2 == 0 {
		return false
	}
	max := int64(math.Sqrt(float64(v)))
	for i := int64(3); i <= max; i += 2 {
		if v%i == 0 {
			return false
		}
	}

	return true
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func isPalindrome(v int64) bool {
	str := strconv.Itoa(int(v))
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}
