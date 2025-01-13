package client 

import (
	"context"
	pb "kasyap1234/book-review-recommendation/proto"

)

func(c*Client)Register(username, password string)(*pb.UserResponse,error){
	req  :=&pb.UserRequest{
		Username: username,
		Password: password,
	}
	return c.userClient.RegisterUser(context.Background(),req)
}

func (c*Client)Login(username, password string)(string,error){
	req :=&pb.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err :=c.userClient.LoginUser(context.Background(),req); 
	if err !=nil {
		return "",err
	}
	return resp.Token,nil 
}





