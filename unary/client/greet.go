package main

import (
	"context"
	"log"

	pb "lab.vps.tips/belin/unary/unary/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Domain: "Belin",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	speed := res.Speed.AsDuration()

	log.Println("Domains:", res.AllDomains, "\n Speed", speed)
}
