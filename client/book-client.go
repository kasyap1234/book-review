package client

import (
	pb "github.com/kasyap1234/book-review/proto"
	"google.golang.org/grpc/metadata"

	"context"
)


func (c *Client) GetBooks() (*pb.GetBooksResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("authorization", c.token))

	req := &pb.GetBooksRequest{}
	return c.bookClient.GetBooks(ctx, req)

}

func (c *Client) AddBook(title, author string) (*pb.BookResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(),
		metadata.Pairs("authorization", c.token))

	req := &pb.AddBookRequest{
		Title:  title,
		Author: author,
	}

	return c.bookClient.AddBook(ctx, req)

}
