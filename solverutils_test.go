package main

import (
	"context"
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/solver/proto"
)

func TestSolve2(t *testing.T) {
	s := Init()
	sol := s.solve(2, 100, 0)

	if sol != 2+8+34 {
		t.Errorf("Bad result: %v", sol)
	}
}

func TestSolve3(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true

	s.friends = append(s.friends, &pbd.RegistryEntry{})
	s.friends = append(s.friends, &pbd.RegistryEntry{})

	sol, err := s.Solve(context.Background(), &pb.SolveRequest{Problem: 3, KeyStart: 1, KeyEnd: 13194, Goal: 13195})
	if err != nil {
		t.Fatalf("Solve fail: %v", err)
	}

	if sol.Solution != int64(29) {
		t.Errorf("Wrong solution: %v", sol)
	}
}

func TestSolve7(t *testing.T) {
	s := Init()
	sol := s.solve(7, 6, 0)

	if sol != 13 {
		t.Errorf("Bad results: %v", sol)
	}
}
