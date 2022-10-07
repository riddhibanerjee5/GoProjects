package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/infogetter/details"
)

const (
	defaultName = "user0"
	defaultID   = "00"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to get")
	id   = flag.String("id", defaultID, "ID to get")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDetailGetterClient(conn)
	allDetails(c)

	/*
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetDetails(ctx, &pb.Request{Name: *name, ID: *id})
		if err != nil {
			log.Fatalf("could not get: %v", err)
		}
		log.Printf("Info: %s", r.GetCreds())
	*/
}

type allRequests struct {
	requests []*pb.Request
}

func (a allRequests) initRequests() []*pb.Request {
	a.requests = []*pb.Request{
		{
			Name: "Riddhi",
			ID:   "05",
		},

		{
			Name: "User2",
			ID:   "02",
		},

		{
			Name: "User3",
			ID:   "03",
		},

		{
			Name: "Bob",
			ID:   "04",
		},
	}

	return a.requests
}

func allDetails(c pb.DetailGetterClient) {
	stream, err := c.GetStreamDetails(context.Background())

	if err != nil {
		log.Fatalf("Error getting stream request: %v", err)
		return
	}

	requests := allRequests{}.initRequests()

	waitResponse := make(chan struct{})

	go func() {
		for _, req := range requests {
			stream.Send(req)

			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error when receiving response: %v", err)
			}

			fmt.Println("Server response: ", res)
		}

		close(waitResponse)
	}()
	<-waitResponse

}

// go run client/main.go --name=Riddhi --id=05
// mess around with GO111MODULE for "cannot find package in" error to go away
