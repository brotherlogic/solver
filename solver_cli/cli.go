package main

import (
	"fmt"
	"log"
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
	res, err := client.Solve(ctx, &pb.SolveRequest{Problem: 1, KeyStart: 1, KeyEnd: 1000})
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	fmt.Printf("Answer: %v (%v) from %v nodes.\n", res.Solution, time.Now().Sub(t), res.Nodes)
}
