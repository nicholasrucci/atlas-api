package route

import (
	"net/http"

	"atlas-api/controllers"
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

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/api/authenticate",
		controllers.Authenticate,
	},
	Route{
		"CreateUser",
		"POST",
		"/api/users/new",
		controllers.CreateUser,
	},
	Route{
		"CreateOrganizaiton",
		"POST",
		"/api/organization/new",
		controllers.CreateOrganization,
	},
}
