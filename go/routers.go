/*
 * Payments API Template
 *
 * A sample openAPI specification that uses a simple user group system as an example to generate a skelton API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"UsersByIdDelete",
		strings.ToUpper("Delete"),
		"/v1/users/{id}",
		UsersByIdDelete,
	},

	Route{
		"UsersByIdGet",
		strings.ToUpper("Get"),
		"/v1/users/{id}",
		UsersByIdGet,
	},

	Route{
		"UsersGet",
		strings.ToUpper("Get"),
		"/v1/users",
		UsersGet,
	},

	Route{
		"UsersPost",
		strings.ToUpper("Post"),
		"/v1/users",
		UsersPost,
	},
}
