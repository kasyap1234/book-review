package client

import (
	"context"
	pb "kasyap1234/book-review-recommendation/proto"
) 


func (c *Client)GetBooks()(*pb.GetBooksResponse,error){
req := &pb.GetBooksRequest{}
return c.bookClient.GetBooks(context.Background(),req)

}

func(c*Client)AddBook(title,author string)(*pb.BookResponse,error){
	req :=&pb.AddBookRequest{
		Title: title,
		Author: author,
	}

	return c.bookClient.AddBook(context.Background(),req); 


}

