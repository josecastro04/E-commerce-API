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
	{
		URI:           "/update_password",
		Method:        http.MethodPut,
		Func:          controllers.UpdatePassword,
		Authorization: true,
	},
	{
		URI:           "/delete_user",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteUser,
		Authorization: true,
	},
}
