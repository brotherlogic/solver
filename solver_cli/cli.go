package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/solver/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	host, port, err := utils.Resolve("solver")
	if err != nil {
		log.Fatalf("Unable to reach solver: %v", err)
	}
	conn, err := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}

	client := pb.NewSolverServiceClient(conn)
	ctx, cancel := utils.ManualContext("solver-cli", "solver", time.Minute)
	defer cancel()

	t := time.Now()
	problem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Bad problem")
	}

	var res *pb.SolveResponse
	switch problem {
	case 1:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 1, KeyStart: 1, KeyEnd: 1000})
	case 2:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 2, KeyStart: 1, KeyEnd: 4000000})
	case 3:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 3, KeyStart: 1, KeyEnd: int64(math.Sqrt(float64(600851475143))), Goal: 600851475143})
	case 4:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 4, KeyStart: 1001, KeyEnd: 999999, Goal: 1000})
	case 5:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 5, KeyStart: 1, KeyEnd: 1 * 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10 * 11 * 12 * 13 * 14 * 15 * 16 * 17 * 18 * 19 * 20, Goal: 20})
	case 6:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 6, KeyStart: 1, KeyEnd: 101})
		fmt.Printf("Answer: %v (%v) from %v nodes.\n", 5050*5050-res.Solution, time.Now().Sub(t), res.Nodes)
	case 7:
		res, err = client.Solve(ctx, &pb.SolveRequest{Problem: 7, KeyStart: 1, KeyEnd: 10001})
	}
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	fmt.Printf("Answer: %v (%v) from %v nodes.\n", res.Solution, time.Now().Sub(t), res.Nodes)
}
