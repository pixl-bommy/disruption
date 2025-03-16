package api

import (
	"fmt"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: CreateEvent called")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/v1: GetEvents called")

	// always return empty array ("stub" for now)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[]"))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: DeleteEvent called")
}
