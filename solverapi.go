package main

import (
	"fmt"

	pb "github.com/brotherlogic/solver/proto"

	"golang.org/x/net/context"
)

// Solve solves a problem
func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
