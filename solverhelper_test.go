package main

import "testing"

func TestIsPrime(t *testing.T) {
	if isPrime(int64(35)) {
		t.Errorf("35 has been reported prime")
	}
}
