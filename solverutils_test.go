package main

import "testing"

func TestSolve2(t *testing.T) {
	s := Init()
	sol := s.solve(2, 100)

	if sol != 2+8+34 {
		t.Errorf("Bad result: %v", sol)
	}
}
