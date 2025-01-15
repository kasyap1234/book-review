package book

import (
	pb "kasyap1234/book-review-recommendation/proto"
	
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer1() {
	listener, err := net.Listen("tcp", ":50051")
if err !=nil {
	log.Fatalln("unable to listen to port"); 

}
log.Println("server started on port :50051")

grpcServer :=grpc.NewServer()
pb.RegisterBookServiceServer(grpcServer,&server{})
if err :=grpcServer.Serve(listener); err !=nil {
	log.Fatalf("Failed to listen to port"); 
}

}
