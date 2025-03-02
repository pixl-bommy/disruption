package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pixl-bommy/disruption/configs"
	"github.com/pixl-bommy/disruption/internal/services"
)

func CreateDisruption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/v1: CreateDisruption called")

	// Read the request json body
	var unsecureJson map[string]string

	err := json.NewDecoder(r.Body).Decode(&unsecureJson)
	if err != nil {
		fmt.Println("api/v1: CreateDisruption unmarshaling request to json failed")
		fmt.Println("   ", err)

		http.Error(w, "request body is not valid json", http.StatusUnprocessableEntity)
		return
	}

	// Extract the name and description from the request json body
	name := unsecureJson["name"]
	description := unsecureJson["description"]

	// Extract the user id from the request json body
	// TODO: should be extracted from the JWT token or some other secure source
	userId := configs.OneAndOnlyUserUid

	// Validate the request json body
	if name == "" || description == "" || userId == "" {
		fmt.Println("api/v1: CreateDisruption invalid request json body")

		http.Error(w, "request body is not valid", http.StatusUnprocessableEntity)
		return
	}

	// Create a new disruption
	disruptionId, err := services.Disruptions.Create(name, description, userId)
	if err != nil {
		fmt.Println("Failed to create disruption:", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Created disruption with id:", disruptionId)

	w.Header().Add("Location", "/disruptions/"+disruptionId)
	w.WriteHeader(http.StatusCreated)
}

func GetDisruptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/v1: STUB: GetDisruptions called")
}

func DeleteDisruption(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api/v1: DeleteDisruption called")

	// get itemId from url
	rawDisruptionId := mux.Vars(r)["disruption-id"]
	_, err := uuid.Parse(rawDisruptionId)
	if err != nil {
		fmt.Println("api/v1: DeleteDisruption invalid disruption id")
		fmt.Println("   ", err)

		http.Error(w, "disruption id is not valid", http.StatusUnprocessableEntity)
		return
	}

	// Extract the user id from the request json body
	// TODO: should be extracted from the JWT token or some other secure source
	userId := configs.OneAndOnlyUserUid

	// Validate the request json body
	if rawDisruptionId == "" || userId == "" {
		fmt.Println("api/v1: DeleteDisruption invalid request")

		http.Error(w, "request is not valid", http.StatusUnprocessableEntity)
		return
	}

	// Create a new disruption
	isDeleted, err := services.Disruptions.Delete(rawDisruptionId, userId)
	if err != nil {
		fmt.Println("api/v1: DeleteDisruption failed to delete:", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isDeleted {
		fmt.Println("api/v1: DeleteDisruption has nothing to delete:", err)

		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println("api/v1: DeleteDisruption deleted disruption with id:", rawDisruptionId)

	w.WriteHeader(http.StatusNoContent)
}
