package main

import (
	"context"
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/solver/proto"
)

func TestSolve4(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true

	s.friends = append(s.friends, &pbd.RegistryEntry{})
	s.friends = append(s.friends, &pbd.RegistryEntry{})

	sol, err := s.Solve(context.Background(), &pb.SolveRequest{Problem: 4, KeyStart: 101, KeyEnd: 9999, Goal: 100})
	if err != nil {
		t.Fatalf("Solve fail: %v", err)
	}

	if sol.Solution != int64(9009) {
		t.Errorf("Wrong solution: %v", sol)
	}
}
