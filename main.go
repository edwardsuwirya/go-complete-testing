package main

import (
	"enigmacamp.com/completetesting/delivery"
	"log"
	"os"
)

func main() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	err := delivery.NewServer(host, port).StartEngine()
	if err != nil {
		log.Fatal(err)
	}
}
