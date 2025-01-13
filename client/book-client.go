package client

import (
	"context"
	pb "kasyap1234/book-review-recommendation/proto"

	"google.golang.org/grpc/internal/metadata"
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
