package client

import (
	"context"
	pb "kasyap1234/book-review-recommendation/proto"

	"google.golang.org/grpc/internal/metadata"
)

func (c *Client) AddReview(book_id, review string, rating int) (*pb.AddReviewResponse, error) {
	md := metadata.Pairs("authorization", "Bearer "+c.token)
    ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &pb.AddReviewRequest{
		BookId: book_id,
		
		Review: review,
		Rating: int32(rating),
	}
	return c.reviewClient.AddReview(ctx, req)
}

func (c *Client) GetReviewByBook(book_id string) (*pb.GetReviewByBookResponse, error) {
	ctx := metadata.NewOutgoingContext(context.Background(), 
    metadata.Pairs("authorization", c.token))

	req := &pb.GetReviewByBookRequest{
		BookId: book_id,
	}
	return c.reviewClient.GetReviewByBook(ctx, req)
}
