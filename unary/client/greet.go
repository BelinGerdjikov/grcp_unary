package main

import (
	"context"
	"log"

	pb "lab.vps.tips/belin/unary/unary/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
