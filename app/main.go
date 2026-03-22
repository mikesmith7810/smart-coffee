package main

import (
	"log"
	"smart-coffee/router"
)

func main() {
	r := router.New()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
