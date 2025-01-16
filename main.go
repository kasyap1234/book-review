package main

import (
	"kasyap1234/book-review-recommendation/client"
	"log"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		log.Fatal("client initializing error:", err)
	}
	defer c.Close()

	// Add a timeout to the context
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	registerResponse, err := c.Register("kasyap", "kasyap1", "password")
	if err != nil {
		log.Printf("Register error in register response: %v", err)
	} else {
		log.Println("Register response:", registerResponse)
	}

	loginResponse, err := c.Login("kasyap1", "password")
	if err != nil {
		log.Printf("Login error: %v", err)
	} else {
		log.Println("Login successful:", loginResponse)
	}
		// Now proceed with book operations
		bookResp, err := c.AddBook("harry potter", "Jk Rowling")
		if err != nil {
			log.Printf("AddBook error: %v", err)
		} else {
			log.Println("AddBook response:", bookResp)
		}
	
}
