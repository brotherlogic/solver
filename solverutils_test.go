package main

import "testing"

func TestSolve2(t *testing.T) {
	s := Init()
	sol := s.solve(2, 100)

	if sol != 1+2+3+5+8+13+21+34+55+89 {
		t.Errorf("Bad result: %v", sol)
	}
}
