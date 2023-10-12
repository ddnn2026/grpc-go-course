package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "David",
	})
	if err != nil {
		log.Fatal("not greet : %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
