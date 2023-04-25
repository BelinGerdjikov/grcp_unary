package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	pb "lab.vps.tips/belin/unary/unary/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	start := time.Now()

	var domains []*pb.Domain

	domain := &pb.Domain{
		Name: "belin", Times: "22",
	}

	domains = append(domains, domain)

	end := time.Since(start)

	speed := durationpb.New(end)

	report := pb.GreetResponse{
		AllDomains: domains, Speed: speed,
	}

	return &report, nil
}
