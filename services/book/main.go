package book

import (
	pb "kasyap1234/book-review-recommendation/proto"
	
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
if err !=nil {
	log.Fatalln("unable to listen to port"); 

}
grpcServer :=grpc.NewServer()
pb.RegisterBookServiceServer(grpcServer,&server{})
if err :=grpcServer.Serve(listener); err !=nil {
	log.Fatalf("Failed to listen to port"); 
}

}
