package route

import (
	"net/http"
)

// Route is a struct contating the name,
// method, pattern, and handler function
// for the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes holds an array of of type Route
type Routes []Route

var routes = Routes{}
