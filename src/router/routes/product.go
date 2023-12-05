package routes

import (
	"api/src/controllers"
	"net/http"
)

var Product = []Route{
	{
		URI:           "/insert_product",
		Method:        http.MethodPost,
		Func:          controllers.InsertProduct,
		Authorization: true,
		OnlyAdmin:     true,
	},
	{
		URI:           "/show_product/{productID}",
		Method:        http.MethodGet,
		Func:          controllers.ShowProduct,
		Authorization: false,
		OnlyAdmin:     false,
	},
}
