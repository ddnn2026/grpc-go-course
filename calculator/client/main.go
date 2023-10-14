package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	doPrimes(c)
}
