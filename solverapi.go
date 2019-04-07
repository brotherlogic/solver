package main

import (
	"fmt"

	pb "github.com/brotherlogic/solver/proto"

	"golang.org/x/net/context"
)

// Solve solves a problem
func (s *Server) Solve(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	if req.Step == 0 {
		return s.distributeProblem(ctx, req)
	}
	return s.solveProblem(req)
}

func (s *Server) solveProblem(req *pb.SolveRequest) (*pb.SolveResponse, error) {
	solution := int64(0)
	for i := req.KeyStart; i < req.KeyEnd; i += req.Step {
		solution += s.solve(req.Problem, i)
	}
	return &pb.SolveResponse{Solution: solution}, nil
}

func (s *Server) distributeProblem(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	solution := int64(0)
	for i := 0; i < len(s.friends); i++ {
		resp, err := s.runSolve(ctx, i, &pb.SolveRequest{
			Problem:  req.Problem,
			KeyStart: int64(i + 1),
			Step:     int64(len(s.friends)),
			KeyEnd:   req.KeyEnd})
		if err != nil {
			return nil, fmt.Errorf("Failed to reach friend %v", s.friends[i].Identifier)
		}
		solution += resp.Solution
	}
	return &pb.SolveResponse{Solution: solution, Nodes: int32(len(s.friends))}, nil
}
