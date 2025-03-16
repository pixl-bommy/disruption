package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

	timestamp := mux.Vars(r)["timestamp"]
	offset := r.URL.Query().Get("offset")

	// make sure offset is defined
	if offset == "" {
		offset = "0"
	}

	fmt.Println("api/v1: GetEvents timestamp (in ms) from url:", timestamp)
	fmt.Println("api/v1: GetEvents offset (in min) from query:", offset)

	// Get the start and end of the day
	start, end, err := getDayLimits(timestamp)
	if err != nil {
		fmt.Println("api/v1: GetEvents failed to parse timestamp:", err)

		// return empty array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}

	// shift the start and end of the day by the offset
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Println("api/v1: GetEvents failed to parse offset:", err)

		// return empty array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}
	start += int64(offsetInt) * 60 * 1000
	end += int64(offsetInt) * 60 * 1000

	fmt.Println("api/v1: GetEvents start of the day:", start)
	fmt.Println("api/v1: GetEvents end of the day:", end)

	// Get the events for the day
	dailyEvents, err := services.DailyEvents.Get(start, end)
	if err != nil {
		fmt.Println("api/v1: GetEvents failed to get events:", err)

		// return empty array
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}

	// always return empty array ("stub" for now)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dailyEvents)
	if err != nil {
		fmt.Println("api/v1: GetEvents failed to encode events:", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("STUB: DeleteEvent called")
}

func getDayLimits(timestamp string) (int64, int64, error) {
	// Parse the timestamp to milliseconds
	milliseconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	// Convert the milliseconds to time
	t := time.UnixMilli(milliseconds)

	// Get the start of the day
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	// Get the end of the day
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)

	return start.UnixMilli(), end.UnixMilli(), nil
}
