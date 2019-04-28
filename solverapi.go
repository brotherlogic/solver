package main

import (
	"fmt"
	"math"

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

	if req.Problem == 5 {
		solution = math.MaxInt64
	}

	if req.Problem == 1 || req.Problem == 3 || req.Problem == 4 || req.Problem == 5 || req.Problem == 6 {
		for i := req.KeyStart; i < req.KeyEnd; i += req.Step {
			if req.Problem == 1 || req.Problem == 6 {
				solution += s.solve(req.Problem, i, req.Goal)
			}
			if req.Problem == 3 || req.Problem == 4 {
				solution = max64(solution, s.solve(req.Problem, i, req.Goal))
			}
			if req.Problem == 5 {
				solution = min64(solution, s.solve(req.Problem, i, req.Goal))
			}
		}
	} else if req.Problem == 2 || req.Problem == 7 {
		solution = s.solve(req.Problem, req.KeyEnd, req.Goal)
	}
	return &pb.SolveResponse{Solution: solution}, nil
}

func (s *Server) asyncSolve(ctx context.Context, i int, req *pb.SolveRequest, c chan int64) {
	resp, err := s.runSolve(ctx, i, &pb.SolveRequest{
		Problem:  req.Problem,
		KeyStart: int64(i + 1),
		Step:     int64(len(s.friends)),
		Goal:     req.Goal,
		KeyEnd:   req.KeyEnd})
	if err != nil {
		c <- -1
	} else {
		c <- resp.Solution
	}
}

func (s *Server) distributeProblem(ctx context.Context, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	solution := int64(0)

	if req.Problem == 5 {
		solution = math.MaxInt64
	}

	if len(s.friends) == 0 {
		return nil, fmt.Errorf("No friends")
	}

	c := make(chan int64)
	currFriends := len(s.friends)
	solutions := 0
	for i := 0; i < len(s.friends); i++ {
		go s.asyncSolve(ctx, i, req, c)
	}

	for solutions < currFriends {
		sol := <-c
		solutions++
		if sol == -1 {
			return nil, fmt.Errorf("Failed to reach a friend")
		}
		if req.Problem == 2 || req.Problem == 7 {
			solution = sol
		} else if req.Problem == 3 || req.Problem == 4 {
			solution = max64(solution, sol)
		} else if req.Problem == 5 {
			solution = min64(solution, sol)
		} else {
			solution += sol
		}
	}
	return &pb.SolveResponse{Solution: solution, Nodes: int32(len(s.friends))}, nil
}
