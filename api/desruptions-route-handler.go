package api

import (
	"fmt"
	"net/http"
)

func CreateDisruption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: CreateDisruption called")
}

func GetDisruptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: GetDisruptions called")
}

func DeleteDisruption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: DeleteDisruption called")
}
