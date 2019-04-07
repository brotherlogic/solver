package main

import (
	"context"
	"testing"

	pbd "github.com/brotherlogic/discovery/proto"
	pb "github.com/brotherlogic/solver/proto"
)

func TestBasic(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true
	s.friends = append(s.friends, &pbd.RegistryEntry{})
	r, err := s.Solve(context.Background(), &pb.SolveRequest{})

	if err != nil {
		t.Errorf("Oops: %v", err)
	}

	if r.Solution == 0 {
		t.Errorf("Problem in solution: %v", r)
	}
}

func TestProblem(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true
	s.friends = append(s.friends, &pbd.RegistryEntry{})
	r, err := s.Solve(context.Background(), &pb.SolveRequest{Step: 10, Problem: 1, KeyStart: 1, KeyEnd: 200})

	if err != nil {
		t.Errorf("Oops: %v", err)
	}

	if r.Solution == 0 {
		t.Errorf("Problem in solution: %v", r)
	}
}

func TestProblem2(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true
	s.friends = append(s.friends, &pbd.RegistryEntry{})
	r, err := s.Solve(context.Background(), &pb.SolveRequest{Step: 10, Problem: 2, KeyStart: 1, KeyEnd: 200})

	if err != nil {
		t.Errorf("Oops: %v", err)
	}

	if r.Solution == 0 {
		t.Errorf("Problem in solution: %v", r)
	}
}

func TestNoFriends(t *testing.T) {
	s := Init()
	s.test = true
	s.pass = true

	r, err := s.Solve(context.Background(), &pb.SolveRequest{Problem: 1, KeyStart: 1, KeyEnd: 200})
	if err == nil {
		t.Errorf("Oops: %v", r)
	}
}

func TestBasicFail(t *testing.T) {
	s := Init()
	s.test = true
	s.friends = append(s.friends, &pbd.RegistryEntry{})
	r, err := s.Solve(context.Background(), &pb.SolveRequest{})

	if err == nil {
		t.Errorf("Oops: %v", r)
	}
}
