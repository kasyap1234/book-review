package client

import (
	"context"
	pb "kasyap1234/book-review-recommendation/proto"
)

func (c *Client) AddReview(book_id, user_id, review string, rating int) (*pb.AddReviewResponse, error) {
	req := &pb.AddReviewRequest{
		BookId: book_id,
		UserId: user_id,
		Review: review,
		Rating: int32(rating),
	}
	return c.reviewClient.AddReview(context.Background(), req)
}

func (c *Client) GetReviewByBook(book_id string) (*pb.GetReviewByBookResponse, error) {
	req := &pb.GetReviewByBookRequest{
		BookId: book_id,
	}
	return c.reviewClient.GetReviewByBook(context.Background(), req)
}
