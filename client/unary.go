package main

import (
	"context"
	"log"
	pb "shyam_kadiyam/grpc/proto"
	"time"
)

func callSayHello(client pb.GrpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	log.Printf("%s", res.Message)
}
