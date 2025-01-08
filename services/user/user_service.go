package user

import (
	"context"
	"fmt"
	pb "kasyap1234/book-review-recommendation/proto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var secretKey = []byte("secret-key")

type Claims struct {
	userID int 

	jwt.Claims
}


type Server struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func generateToken(userID int ) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func VerifyToken(tokenString string) (*Claims,error) {
	claims :=&Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:   uint32(user.ID),
		Name: user.Name,
	}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	var user User

	if err := s.db.First(&user, "username=? AND password= ? ", req.Username, req.Password).Error; err != nil {
		return nil, err
	}
	
	
	tokenString, err := generateToken(int(user.ID))
	
	if err !=nil {
		fmt.Println("unable to generate jwt token"); 
	}
return &pb.TokenResponse{
	Token: tokenString,
},nil 

}


