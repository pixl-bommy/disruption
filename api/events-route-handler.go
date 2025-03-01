package api

import (
	"fmt"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: CreateEvent called")
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: GetEvents called")
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: DeleteEvent called")
}
