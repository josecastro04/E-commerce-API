package routes

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI           string
	Method        string
	Func          func(w http.ResponseWriter, r *http.Request)
	Authorization bool
	OnlyAdmin     bool
}

func Config(router *mux.Router) *mux.Router {
	routes := User
	routes = append(routes, Login)
	routes = append(routes, Product...)

	for _, route := range routes {
		if route.Authorization {
			if route.OnlyAdmin {
				router.HandleFunc(route.URI, middlewares.Authenticate(middlewares.Authorize("admin", route.Func))).Methods(route.Method)
			} else {
				router.HandleFunc(route.URI, middlewares.Authenticate(route.Func)).Methods(route.Method)
			}
		} else {
			router.HandleFunc(route.URI, route.Func).Methods(route.Method)
		}
	}

	return router
}
