package main

import (
	"context"
	"flag"
	"fmt"
	pbproto "grpc-network-speed/pkg/proto-pb"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

// Ref: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/client/client.go

var (
	serverAddr = flag.String("server_addr", "172.27.172.134:2000", "The server address in the format of host:port")
)

// Invoking rpc call towards the server
func askGetNAme(ctx context.Context, client pbproto.StudentServiceClient) int64 {
	fmt.Println("\n>>>>>>>>>>>>>>.sendGetMarkReq")
	req := new(pbproto.Request)
	req.JobId = "Job-ID-123"
	req.SlNo = 123
	btes, _ := proto.Marshal(req)
	fmt.Printf("Marshalled Data : %v\n", btes)
	fmt.Printf("Get-mark-for-Prince : %v\n", req)
	startTime := time.Now().UnixNano()
	resp, err := client.GetName(ctx, req)
	endTime := time.Now().UnixNano()
	timeDiff := endTime - startTime
	if err != nil {
		log.Fatalf("could not get name: %v\n", err)
	}

	log.Printf("\n\n========================Time Difference=========\nResponse : %v\nStart Time: %v\nEnd Time : %v\nTime Diff : %v nano seconds\n==================\n\n", resp, startTime, endTime, timeDiff)
	return timeDiff
}

func main() {
	log.Printf("Initialised....  \n")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(10*time.Second))

	// Establishing the server connection
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("COnnection established...\n")

	defer conn.Close()

	// Constructing a client object
	client := pbproto.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if ctx == nil || client == nil {
		fmt.Println("Conext or client is nil")
	}

	var totalTimeDiff int64
	for i := 1; i <= 10; i++ {
		totalTimeDiff = totalTimeDiff + askGetNAme(ctx, client)
	}
	average := totalTimeDiff / 10
	oneDirection := average / 2
	var oneDirectionMs float64
	oneDirectionMs = float64(oneDirection) / 1000000
	// fmt.Printf("\n\n========== Final Result after 10 run==========\n\nAverage Time : %v\nTime for one direction (Nano) : %v\nTime for one direction (Milli) : %v\n===========\n\n", average, oneDirection, oneDirectionMs)
	log.Printf("\n\n========== Final Result after 10 run==========\n\nAverage Time : %v\nTime for one direction (Nano) : %v\nTime for one direction (Milli) : %v\n===========\n\n", average, oneDirection, oneDirectionMs)
	// sendGetAllMarksReq(ctx, client)
}
