package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	SkipAuth    bool
}

var routes = Routes{
	Route{
		Name:        "RegisterCompany",
		Method:      "POST",
		Pattern:     "/companies/register",
		HandlerFunc: RegisterCompany,
	},
	Route{
		Name:        "GetCompany",
		Method:      "GET",
		Pattern:     "/companies/{id}",
		HandlerFunc: GetCompany,
	},
	Route{
		Name:        "GetAllContracts",
		Method:      "GET",
		Pattern:     "/contracts",
		HandlerFunc: GetAllContracts,
	},
	// Route{
	// 	Name:        "RegisterUser",
	// 	Method:      "POST",
	// 	Pattern:     "/users/register",
	// 	HandlerFunc: RegisterUser,
	// },
}

func NewRouter() *mux.Router {
	s := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router := s.StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
