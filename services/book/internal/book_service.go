package bookservice 

import (
	"context"
	
pb "github.com/kasyap1234/book-review/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedBookServiceServer
	db *gorm.DB
}


func (s *server) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	userId, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}
	var books []Book
	if err := s.db.Where("userID=?", userId).Find(&books).Error; err != nil {
		return nil, status.Error(codes.NotFound, "user does not exist")
	}
	var pbBooks []*pb.Book
	for _, book := range books {
		pbBooks = append(pbBooks, &pb.Book{
			Id:     uint32(book.ID),
			Title:  book.Title,
			Author: book.Author,
		})
	}

	return &pb.GetBooksResponse{Books: pbBooks}, nil
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.BookResponse, error) {
	userId, ok := ctx.Value("userID").(int)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	book := Book{
		Title:  req.Title,
		Author: req.Author,
		UserID: userId,
	}

	if err := s.db.Create(&book).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add book: %v", err)
	}

	return &pb.BookResponse{
		Id:     uint32(book.ID),
		Title:  book.Title,
		Author: book.Author,
	}, nil
}
