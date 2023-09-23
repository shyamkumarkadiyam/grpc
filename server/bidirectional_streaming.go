package main

import (
	"io"
	"log"
	pb "shyam_kadiyam/grpc/proto"
)

func (bs *helloServer) SayHelloBidirectionalStream(stream pb.GrpcService_SayHelloBidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
