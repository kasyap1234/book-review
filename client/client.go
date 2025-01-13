package client

import (
	
	pb "kasyap1234/book-review-recommendation/proto"

	"google.golang.org/grpc"
	
)

type Client struct {
	bookClient   pb.BookServiceClient
	reviewClient pb.ReviewServiceClient
	userClient   pb.UserServiceClient
	bookConn     *grpc.ClientConn
	userConn     *grpc.ClientConn
	reviewConn   *grpc.ClientConn
	token string 
}

func NewClient() (*Client, error) {
	bookconn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	userConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	reviewConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		bookClient:   pb.NewBookServiceClient(bookconn),
		reviewClient: pb.NewReviewServiceClient(reviewConn),
		userClient:   pb.NewUserServiceClient(userConn),
		bookConn:     bookconn,
		reviewConn:   reviewConn,
		userConn:     userConn,
	}, nil

}

func (c *Client) Close() {
	c.userConn.Close()
	c.bookConn.Close()
	c.reviewConn.Close()
}
