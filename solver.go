package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pbd "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
	"github.com/brotherlogic/keystore/client"
	pb "github.com/brotherlogic/solver/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	friends []*pbd.RegistryEntry
	test    bool
	pass    bool
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		make([]*pbd.RegistryEntry, 0),
		false,
		false,
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterSolverServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{
		&pbg.State{Key: "friends", Value: int64(len(s.friends))},
	}
}

func (s *Server) runSolve(ctx context.Context, index int, req *pb.SolveRequest) (*pb.SolveResponse, error) {
	if s.test {
		if s.pass {
			return s.Solve(ctx, req)
		}
		return nil, fmt.Errorf("Built to fail")
	}

	conn, err := s.DoDial(s.friends[index])
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewSolverServiceClient(conn)
	return client.Solve(ctx, req)
}

func (s *Server) findFriends(ctx context.Context) error {
	entries, err := utils.ResolveAll("solver")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Identifier != s.Registry.Identifier {
			found := false
			for _, fr := range s.friends {
				if entry.Identifier == fr.Identifier {
					found = true
				}
			}

			if !found {
				s.friends = append(s.friends, entry)
			}
		}
	}

	return nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	server.PrepServer()
	server.Register = server

	server.RegisterServer("solver", false)

	server.RegisterRepeatingTask(server.findFriends, "friends", time.Minute*5)

	server.SendTrace = false

	fmt.Printf("%v\n", server.Serve())
}
