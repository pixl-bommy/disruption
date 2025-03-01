package api

import (
	"fmt"
	"net/http"
)

// NOTE: this is a stub implementation
// Later on, we will replace this with a real implementation
// that fetches the user's information from a database. And
// we will also add authentication to make sure only the
// user who is logged in can access restricted resources.

const (
	// the one and only user
	userName = "Bommy"
	uid      = "58032f45-80c4-4fdf-8868-d4608ba37b55"
)

func GetUserMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: GetUserMe called")

	// return a JSON response containing the user's information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"uid": "%s", "name": "%s"}`, uid, userName)
}
