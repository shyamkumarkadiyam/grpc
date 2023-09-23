package main

import (
	"io"
	"log"
	pb "shyam_kadiyam/grpc/proto"
)

func (cl *helloServer) SayHelloClientStream(stream pb.GrpcService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
