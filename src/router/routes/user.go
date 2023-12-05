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
		OnlyAdmin:     false,
	},
	{
		URI:           "/show_userinfo",
		Method:        http.MethodGet,
		Func:          controllers.ShowUserInfo,
		Authorization: true,
		OnlyAdmin:     false,
	},
	{
		URI:           "/update_password",
		Method:        http.MethodPut,
		Func:          controllers.UpdatePassword,
		Authorization: true,
		OnlyAdmin:     false,
	},
	{
		URI:           "/delete_user",
		Method:        http.MethodDelete,
		Func:          controllers.DeleteUser,
		Authorization: true,
		OnlyAdmin:     false,
	},
}
