package client 

import (
	"context"
	pb "github.com/kasyap1234/book-review/proto"

)



func(c*Client)Register(name,username, password string)(*pb.UserResponse,error){
	req  :=&pb.UserRequest{
		Name: name,
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
	c.token=resp.Token
	return resp.Token,nil
}






