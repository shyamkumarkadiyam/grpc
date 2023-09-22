package main

import (
	"log"
	pb "shyam_kadiyam/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error connecting %v", err)

	}
	defer conn.Close()

	client := pb.NewGrpcServiceClient(conn)

	//names := &pb.NamesList{
	//	Name: []string{"Shyam", "Kumar", "Kadiyam"},
	//}

	callSayHello(client)
}
