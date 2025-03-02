package api

import (
	"fmt"
	"net/http"

	"github.com/pixl-bommy/disruption/configs"
)

func GetUserMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: GetUserMe called")

	// return a JSON response containing the user's information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"uid": "%s", "name": "%s"}`, configs.OneAndOnlyUserUid, configs.OneAndOnlyUserName)
}
