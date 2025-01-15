package review

import (
	pb "kasyap1234/book-review-recommendation/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer2(){
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln("Failed to listen to port")

	}
	log.Println("server started on port 50052")
	grpcServer := grpc.NewServer()
	pb.RegisterReviewServiceServer(grpcServer, &server{})
	log.Println("Review service running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("err listening ")
	}

}
