package main

import "testing"

func TestIsPrime(t *testing.T) {
	if isPrime(int64(35)) {
		t.Errorf("35 has been reported prime")
	}
}

var palindromeData = []struct {
	value int64
	palin bool
}{
	{int64(123), false},
	{int64(121), true},
	{int64(1321), false},
	{int64(1221), true},
}

func TestPalindrome(t *testing.T) {
	for _, data := range palindromeData {
		if isPalindrome(data.value) != data.palin {
			t.Errorf("Failed: %v is reported %v (should be %v)", data.value, isPalindrome(data.value), data.palin)
		}
	}
}
