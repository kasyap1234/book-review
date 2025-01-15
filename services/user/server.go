package user

import (
	pb "kasyap1234/book-review-recommendation/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)



func StartServer3() {
	 s := &Server{}
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalln("failed to listen to port ")
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer,s); 
	log.Println("service running on port 50053")
	if err= grpcServer.Serve(listener); err !=nil {
		log.Fatal("failed to listen properly ")
	}
}

