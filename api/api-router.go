package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	// set up cors middleware
	router.Use(corsMiddleware)

	// api/v1: desruptions
	router.HandleFunc("/api/v1/disruptions", CreateDisruption).Methods("POST")
	router.HandleFunc("/api/v1/disruptions", GetDisruptions).Methods("GET")
	router.HandleFunc("/api/v1/disruptions/{disruption-id}", DeleteDisruption).Methods("DELETE")

	// api/v1: events
	router.HandleFunc("/api/v1/events", CreateEvent).Methods("POST")
	router.HandleFunc("/api/v1/events/{timestamp}", GetEvents).Methods("GET")
	router.HandleFunc("/api/v1/events/{event-id}", DeleteEvent).Methods("DELETE")

	// api/v1: users
	router.HandleFunc("/api/v1/users/me", GetUserMe).Methods("GET")

	// if everything before doesn't match, fallback to static file server
	var staticFileServer = http.FileServer(http.Dir("./ui/dist"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", staticFileServer))

	return router
}

// WORKAROUND: "corsMiddleware" is used for development purposes to allow cross-origin
// requests to the API server. In production, this should not be used as this server
// serves the UI and API from the same origin.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
