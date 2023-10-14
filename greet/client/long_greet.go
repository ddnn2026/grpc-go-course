package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "David"},
		{FirstName: "Nguyen"},
		{FirstName: "ddnn2026"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receiving response form LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
