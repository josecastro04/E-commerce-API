package routes

import (
	"api/src/controllers"
	"net/http"
)

var Login = Route{
	URI:           "/login",
	Method:        http.MethodPost,
	Func:          controllers.LoginUser,
	Authorization: false,
	OnlyAdmin:     false,
}
