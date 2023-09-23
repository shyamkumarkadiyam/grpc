package main

import (
	"context"
	"io"
	"log"
	pb "shyam_kadiyam/grpc/proto"
)

func callSayHelloServerStream(client pb.GrpcServiceClient, names *pb.NamesList) {

	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}

	for {
		HelloRequest, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not greet:%v", err)
		}
		log.Println(HelloRequest)
	}
	log.Println("end of stream")

}
