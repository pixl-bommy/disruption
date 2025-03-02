package main

import (
	"fmt"
	"net/http"

	"github.com/pixl-bommy/disruption/api"
)

func main() {
	fmt.Println("Starting the server...")

	// register API routes
	http.Handle("/", api.NewRouter())

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
