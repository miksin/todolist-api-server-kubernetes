package main

import (
	"net/http"

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
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(logger(route.HandlerFunc, route.Name))
	}
	return router
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	Route{
		Name:        "TodoIndex",
		Method:      "GET",
		Pattern:     "/todos",
		HandlerFunc: TodoIndex,
	},
	Route{
		Name:        "TodoShow",
		Method:      "GET",
		Pattern:     "/todos/{todoId}",
		HandlerFunc: TodoShow,
	},
	Route{
		Name:        "TodoCreate",
		Method:      "POST",
		Pattern:     "/todos",
		HandlerFunc: TodoCreate,
	},
	Route{
		Name:        "TodoDelete",
		Method:      "DELETE",
		Pattern:     "/todos/{todoId}",
		HandlerFunc: TodoDelete,
	},
}
