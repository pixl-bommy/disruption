package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pixl-bommy/disruption/configs"
	"github.com/pixl-bommy/disruption/internal/services"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/v1: CreateEvent called")

	// Read the request json body
	var unsecureJson map[string]string

	err := json.NewDecoder(r.Body).Decode(&unsecureJson)
	if err != nil {
		fmt.Println("api/v1: CreateEvent unmarshaling request to json failed")
		fmt.Println("   ", err)

		http.Error(w, "request body is not valid json", http.StatusUnprocessableEntity)
		return
	}

	// Extract the name and description from the request json body
	disruptionId := unsecureJson["disruptionId"]

	// Extract the user id from the request json body
	// TODO: should be extracted from the JWT token or some other secure source
	userId := configs.OneAndOnlyUserUid

	// Validate the request json body
	if disruptionId == "" || userId == "" {
		fmt.Println("api/v1: CreateEvent invalid request json body")

		http.Error(w, "request body is not valid", http.StatusUnprocessableEntity)
		return
	}

	// Create a new event
	eventId, err := services.DailyEvents.Create(disruptionId, userId)
	if err != nil {
		fmt.Println("api/v1: CreateEvent failed to create event:", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("api/v1: CreateEvent succeeded - id:", eventId)

	// TODO: implement route for -> w.Header().Add("Location", "/events/"+eventId)
	w.WriteHeader(http.StatusCreated)
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
