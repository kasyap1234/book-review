package main 


import (
 "kasyap1234/book-review-recommendation/client"
"context"
"log"
)


func main(){
	c,err :=client.NewClient(); 
	if err !=nil {
		log.Fatal(err); 
	}

	defer c.Close(); 
	registerResponse,err :=c.Register("kasyap1","password")
log.Println(registerResponse)
bookResp, err :=c.AddBook("harry potter","Jk Rowling"); 
log.Println(bookResp)

}
