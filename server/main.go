package main

import (
	"log"
	"net"
	pb "shyam_kadiyam/grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct {
	pb.GrpcServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error starting server %v", err)
	}
	//NewServer creates a gRPC server which has no service registered and has not started to accept requests yet.
	grpcServer := grpc.NewServer()
	//Registering gprcServer with helloServer struct
	pb.RegisterGrpcServiceServer(grpcServer, &helloServer{})
	//Serve accepts incoming connections on the listener lis
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("error starting grpc server %v", err)
	}

}
