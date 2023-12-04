package routes

import (
	"api/src/controllers"
	"net/http"
)

var User = []Route{
	{
		URI:           "/create_user",
		Method:        http.MethodPost,
		Func:          controllers.CreateUser,
		Authorization: false,
	},
	{
		URI:           "/show_userinfo",
		Method:        http.MethodGet,
		Func:          controllers.ShowUserInfo,
		Authorization: true,
	},
}
