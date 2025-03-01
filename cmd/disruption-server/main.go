package main

import (
	"fmt"
	"net/http"

	"github.com/pixl-bommy/disruption/api"
	"github.com/pixl-bommy/disruption/internal/storage"
)

func main() {
	fmt.Println("Starting the server...")

	// register API routes
	http.Handle("/", api.NewRouter())

	// stub to test redis connection
	context := storage.GetContext()
	response, err := storage.Redis().Ping(*context).Result()
	if err != nil {
		fmt.Println("Failed to connect to redis:", err)
		return
	}
	fmt.Println("Connected to redis:", response)

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
