package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter is creating the router using that
// is being used throughout the project using
// gorilla mux
func NewRouter() *mux.Router {
	// create new router
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// Serves files in public directory at the root url
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))

	return router

}

// JSONHandler is being put in all of the handler functions to
// make the Content Type of the responses as JSON.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
