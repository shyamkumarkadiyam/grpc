package main

import (
	"log"
	pb "shyam_kadiyam/grpc/proto"
	"time"
)

func (s *helloServer) SayHelloServerStream(req *pb.NamesList, stream pb.GrpcService_SayHelloServerStreamServer) error {
	log.Printf("got request%v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
