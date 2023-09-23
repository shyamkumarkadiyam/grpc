package main

import (
	"context"
	"log"
	pb "shyam_kadiyam/grpc/proto"
	"time"
)

func callSayHelloClientStream(client pb.GrpcServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Name {
		req := &pb.HelloReq{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Message)
}
