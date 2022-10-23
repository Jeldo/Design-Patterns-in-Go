package main

import "fmt"

var client *Client

type Client struct{}

func GetClient() *Client {
	if client == nil {
		client = &Client{}
		fmt.Println("Creating a new client")
	}
	fmt.Println("Returning the existing client")
	return client
}

func main() {
	// Get the only instance of the singleton
	for i := 0; i < 10; i++ {
		_ = GetClient()
	}
}
