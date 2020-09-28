package main

import (
	"context"
	"flag"
	"fmt"
	pbproto "grpc-network-speed/pkg/proto-pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Ref: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go

var (
	port = flag.Int("port", 2000, "The server port")
)

// Server struct for invoking server calls
type Server struct {
}

// GetName refers to the method getting called for RPC call invoked for Get Single Subject Mark.
func (server *Server) GetName(ctx context.Context, req *pbproto.Request) (*pbproto.Response, error) {
	fmt.Printf("Request received for GetMark() request. Req : %v \n", req)
	resp := new(pbproto.Response)
	resp.JobId = req.JobId
	resp.Name = "Prince Pereira"
	resp.SlNo = req.SlNo

	fmt.Printf("Response send for GetMark() request. Req : %v , Resp : %v\n", req, resp)
	return resp, nil
}

func main() {
	flag.Parse()
	log.Printf("Started Server...")
	// Listening to the port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Started listening...")
	// Constructing a grpc server object.
	grpcServer := grpc.NewServer()

	pbproto.RegisterStudentServiceServer(grpcServer, &Server{})
	log.Printf("Server registered...")
	// Starting the server
	grpcServer.Serve(lis)
}
