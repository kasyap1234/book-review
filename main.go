package main

import (
	"kasyap1234/book-review-recommendation/client"

	"log"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	registerResponse, err := c.Register("kasyap1", "password")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(registerResponse)
	bookResp, err := c.AddBook("harry potter", "Jk Rowling")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(bookResp)

}
