package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/infogetter/details"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDetailGetterServer
}

func (s *server) GetDetails(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: Name - %v, ID - %v", in.GetName(), in.GetID())
	resp := &pb.Response{Creds: "Credentials: Name: " + in.GetName() + " ID: " + in.GetID()}

	return resp, nil
}

func (s *server) GetStreamDetails(stream pb.DetailGetter_GetStreamDetailsServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error from client stream request: %v", err)
		}

		reqID, reqName := req.GetID(), req.GetName()
		fmt.Printf("Request: Name - %v, ID - %v \n", reqName, reqID)

		res := stream.Send(&pb.Response{
			Creds: fmt.Sprintf("Credentials: Name: %v, ID: %v", reqName, reqID),
		})

		if res != nil {
			log.Fatalf("Error from server stream response: %v", res)
		}

	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDetailGetterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
