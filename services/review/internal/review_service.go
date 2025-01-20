package reviewservice 

import (
	"context"
	"errors"
	rb "github.com/kasyap1234/book-review/proto"

	"strconv"

	"gorm.io/gorm"
)

type server struct {
	rb.UnimplementedReviewServiceServer
	db *gorm.DB
}

type Review struct {
	ID     int `gorm:"primaryKey"`
	UserID int `gorm:"not null"`
	BookID int `gorm:"not null"`
	Review string
	Rating int `gorm:"not null"`
}

func NewUserServer(db *gorm.DB) *server {
    return &server{
        db: db,
    }
}

func (r *server) AddReview(ctx context.Context, req *rb.AddReviewRequest) (*rb.AddReviewResponse, error) {
	userId, ok := ctx.Value("userID").(string)
	if !ok || userId != req.UserId {
		return nil, errors.New("unauthorized token ")
	}
	review := Review{
		UserID: parseInt(req.UserId),
		BookID: parseInt(req.BookId),
		Review: req.Review,
		Rating: int(req.Rating),
	}
	if err := r.db.Create(&review).Error; err != nil {
		return nil, err
	}

	return &rb.AddReviewResponse{
		Message: "Successfully added review",
	}, nil

}

func (r *server) GetReviewByBook(ctx context.Context, req *rb.GetReviewByBookRequest) (*rb.GetReviewByBookResponse, error) {
	var reviews []Review
	if err := r.db.Where("book_id = ?", req.BookId).Find(&reviews).Error; err != nil {
		return nil, err
	}

	response := &rb.GetReviewByBookResponse{}
	for _, r := range reviews {
		response.Reviews = append(response.Reviews, &rb.Review{
			UserId: strconv.Itoa(r.UserID),
			BookId: strconv.Itoa(r.BookID),
			Review: r.Review,
			Rating: int32(r.Rating),
		})
	}

	return response, nil
}
func parseInt(number string) int {
	int_number, _ := strconv.Atoi(number)
	return int_number

}
