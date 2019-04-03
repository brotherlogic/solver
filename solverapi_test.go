package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/solver/proto"
)

func TestBasic(t *testing.T) {
	s := Init()
	_, err := s.Solve(context.Background(), &pb.SolveRequest{})
	if err == nil {
		t.Errorf("Should be nil")
	}
}
