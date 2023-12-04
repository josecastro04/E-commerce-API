package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI           string
	Method        string
	Func          func(w http.ResponseWriter, r *http.Request)
	Authorization bool
}

func Config(router *mux.Router) *mux.Router {
	routes := User
	routes = append(routes, Login)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return router
}
